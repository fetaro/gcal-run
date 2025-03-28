package gcal_run

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"google.golang.org/api/calendar/v3"
	"time"
)

type Runner struct {
	Config       *common.Config
	AppDir       string
	EventIDStore *EventIDStore
}

var (
	maxRetryNum = 100
	retrySec    = 30
)

func NewRunner(config *common.Config, appDir string) *Runner {
	return &Runner{
		Config:       config,
		AppDir:       appDir,
		EventIDStore: NewEventIDStore(common.GetEventIDStorePath(appDir)),
	}
}

func (r *Runner) Run() error {

	logger, err := GetLogger(common.GetLogPath(common.GetAppDir()))
	if err != nil {
		return err
	}
	gCalendar := NewCalendar(r.Config.CredentialPath, common.GetTokenPath(r.AppDir))
	var events *calendar.Events

	for i := 1; i <= maxRetryNum; i++ {
		events, err = gCalendar.GetCalendarEvents(time.Now())
		if err != nil {
			if i == maxRetryNum {
				return fmt.Errorf("カレンダーイベントの取得に%d秒間隔で%d回連続で失敗しました。", retrySec, maxRetryNum)
			} else {
				logger.Warn("カレンダーイベントの取得に失敗。%d秒後にリトライします(%d回目/%d回):%v", retrySec, i, maxRetryNum, err)
				time.Sleep(time.Duration(retrySec) * time.Second)
			}
		} else {
			break
		}
	}
	if len(events.Items) == 0 {
		logger.Debug("no gCalendar events")
		return nil
	}
	for _, item := range events.Items {
		event, err := NewEvent(item)
		if err != nil {
			logger.Debug("\"%s\" はオンライン会議ではない", item.Summary)
			continue
		}
		timeToStartSec := event.TimeToStartSec()
		if timeToStartSec < 0 {
			// 開始時刻を過ぎている。現在時刻より後のイベントを取っているため、基本的にはありえない
			logger.Debug("\"%s\" は既に%d分前に開始している", event, -timeToStartSec/60)
			continue
		} else if timeToStartSec < r.Config.MinutesAgo*60 {
			eventAlreadyRun, err := r.EventIDStore.IsInclude(event.ID)
			if err != nil {
				return fmt.Errorf("failed to load event id list from eventIDStore: %v", err)
			}
			if eventAlreadyRun {
				logger.Debug("\"%s\" は既に開始済み", event)
			} else {
				logger.Info("\"%s\" は%d分%d秒後に開始なので、TV会議開始", event, timeToStartSec/60, timeToStartSec%60)
				err := common.OpenUrl(r.Config.BrowserApp, event.URL)
				if err != nil {
					return fmt.Errorf("failed to open event url: %v", err)
				}
				err = r.EventIDStore.SaveID(event.ID)
				if err != nil {
					return fmt.Errorf("failed to save event id to eventIDStore: %v", err)
				}
			}
		} else {
			if timeToStartSec/60 < 5 {
				logger.Info("\"%s\" は%d分%d秒後に開始なので、何もしない", event, timeToStartSec/60, timeToStartSec%60)
			} else {
				logger.Debug("\"%s\" は%d分%d秒後に開始なので、何もしない", event, timeToStartSec/60, timeToStartSec%60)
			}
		}
	}
	return nil
}

func (r *Runner) CleanUp() error {
	logger, err := GetLogger(common.GetLogPath(common.GetAppDir()))
	if err != nil {
		return err
	}
	err = r.EventIDStore.Clear()
	if err != nil {
		return fmt.Errorf("failed to clear event id store: %v", err)
	}
	logger.Debug("clean up event id store")
	return nil
}
