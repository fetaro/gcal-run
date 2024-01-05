package lib

import (
	"fmt"
	"log"
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
	eventIDStore := NewEventIDStore(r.Config.EventIDStorePath)
	calender := NewCalender(r.Config)
	events, err := calender.GetCalenderEvents(time.Now())
	if err != nil {
		return fmt.Errorf("failed to get calender events: %v", err)
	}
	if len(events.Items) == 0 {
		log.Println("カレンダーにイベントがないため、何もしない")
		return nil
	}
	for _, item := range events.Items {
		event, err := NewEvent(item)
		if err != nil {
			continue
		}
		timeToStartSec := event.TimeToStartSec()
		if timeToStartSec < 0 {
			log.Printf("「%s」は開始時間を過ぎている\n", event)
		} else if timeToStartSec < r.Config.SearchMinutes*60 {
			eventAlreadyRun, err := eventIDStore.IsInclude(event.ID)
			if err != nil {
				return fmt.Errorf("failed to load event id list from eventIDStore: %v", err)
			}
			if eventAlreadyRun {
				log.Printf("「%s」は%d分後に開始だが、既に開始済みなので、何もしない\n", event, timeToStartSec/60)
			} else {
				log.Printf("「%s」は%d分後に開始するため、%sのブラウザで会議を開始\n", event, timeToStartSec/60, r.Config.BrowserApp)
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
			log.Printf("「%s」は%d分後に開始なので、何もしない\n", event, timeToStartSec/60)
		}
	}
	return nil
}
