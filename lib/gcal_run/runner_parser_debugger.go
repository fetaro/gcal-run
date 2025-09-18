package gcal_run

import (
	"time"

	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"google.golang.org/api/calendar/v3"
)

func (r *Runner) ParserDebugRun() error {
	// 強制起動に失敗したイベントの原因を探るために、イベントのパースに失敗したイベントを印字する

	logger, err := GetLogger(common.GetLogPath(common.GetAppDir()))
	if err != nil {
		return err
	}
	gCalendar := NewCalendar(r.Config.CredentialPath, common.GetTokenPath(r.AppDir))
	var events *calendar.Events
	daysAgo := 14
	startTime := time.Now().Add(-time.Hour * 24 * time.Duration(daysAgo))
	maxResult := 300
	logger.Info("%d日前から %d 件のイベントを取得", daysAgo, maxResult)
	events, err = gCalendar.GetCalendarEvents(startTime, maxResult)
	if err != nil {
		return err
	}
	if len(events.Items) == 0 {
		logger.Debug("no gCalendar events")
		return nil
	}
	for _, item := range events.Items {
		_, err := NewEvent(item)
		if err != nil {
			// パースに失敗したイベントを印字
			if item.Location == "" {
				continue
			}
			logger.Info("==========================================================================")
			logger.Info("item.Summary: %s", item.Summary)
			logger.Info("item.Location: %s", item.Location)
			logger.Info("item.Description: %s", item.Description)
			logger.Info("err: %v", err)
		}
	}
	return nil
}
