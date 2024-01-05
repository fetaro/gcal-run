package lib

import (
	"fmt"
	"os/exec"
	"time"
)

type Runner struct {
	Config *Config
}

func NewRunner(config *Config) *Runner {
	return &Runner{
		Config: config,
	}
}

func (r *Runner) Run() error {
	logger := GetLogger()
	eventIDStore := NewEventIDStore(r.Config.EventIDStorePath)
	calendar := NewCalendar(r.Config)
	events, err := calendar.GetCalendarEvents(time.Now())
	if err != nil {
		return fmt.Errorf("failed to get calendar events: %v", err)
	}
	if len(events.Items) == 0 {
		logger.Info("no calendar events")
		return nil
	}
	for _, item := range events.Items {
		event, err := NewEvent(item)
		if err != nil {
			logger.Info("何もしない \"%s\" はオンライン会議ではない", item.Summary)
			continue
		}
		timeToStartSec := event.TimeToStartSec()
		if timeToStartSec < 0 {
			// 開始時刻を過ぎている。現在時刻より後のイベントを取っているため、基本的にはありえない
			logger.Info("何もしない \"%s\" は既に%d分前に開始している", event, -timeToStartSec/60)
			continue
		} else if timeToStartSec < r.Config.SearchMinutes*60 {
			eventAlreadyRun, err := eventIDStore.IsInclude(event.ID)
			if err != nil {
				return fmt.Errorf("failed to load event id list from eventIDStore: %v", err)
			}
			if eventAlreadyRun {
				logger.Info("何もしない \"%s\" は既に開始済み", event)
			} else {
				logger.Info("開始　　　 \"%s\" は%d分後に開始", event, timeToStartSec/60)
				err := exec.Command("open", "-a", r.Config.BrowserApp, event.URL).Run()
				if err != nil {
					return fmt.Errorf("failed to open event url: %v", err)
				}
				err = eventIDStore.SaveID(event.ID)
				if err != nil {
					return fmt.Errorf("failed to save event id to eventIDStore: %v", err)
				}
			}
		} else {
			logger.Info("何もしない \"%s\" は%d分後に開始", event, timeToStartSec/60)
		}
	}
	return nil
}
