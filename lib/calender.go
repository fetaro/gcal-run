package lib

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	ApiMaxResult = 4
)

type GCal struct {
	Config *Config
}

func NewCalender(config *Config) *GCal {
	return &GCal{
		Config: config,
	}
}

func (g *GCal) GetCalenderEvents(basisTime time.Time) (*calendar.Events, error) {
	ctx := context.Background()
	b, err := os.ReadFile(g.Config.CredentialPath)
	if err != nil {
		return nil, fmt.Errorf("fail to read credential file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("fail to parse credential file to config: %v", err)
	}

	f, err := os.Open(g.Config.TokenPath)
	if err != nil {
		return nil, fmt.Errorf("fail to open token file: %v", err)
	}
	token := new(oauth2.Token)
	err = gob.NewDecoder(f).Decode(token)
	if err != nil {
		return nil, fmt.Errorf("fail to decode token file: %v", err)
	}
	oAuthClient := config.Client(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("fail to make OAuth client: %v", err)
	}
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(oAuthClient))
	if err != nil {
		return nil, fmt.Errorf("fail to make CalendarService: %v", err)
	}

	log.Printf("Googleカレンダーから、%s以降のイベントを最大%d件取得\n", basisTime.Format(time.RFC3339), ApiMaxResult)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(basisTime.Format(time.RFC3339)).MaxResults(ApiMaxResult).OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("fail to list events: %v", err)
	}
	log.Printf("%d件のイベントを取得\n", len(events.Items))
	return events, nil
}
