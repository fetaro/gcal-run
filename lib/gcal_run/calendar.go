package gcal_run

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
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
	Config *common.Config
}

func NewCalendar(config *common.Config) *GCal {
	return &GCal{
		Config: config,
	}
}

func (g *GCal) GetCalendarEvents(basisTime time.Time) (*calendar.Events, error) {
	logger := GetLogger()
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

	logger.Debug("カレンダーから最大 %d 件のイベントを取得", ApiMaxResult)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(basisTime.Format(time.RFC3339)).MaxResults(ApiMaxResult).OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("fail to list events: %v", err)
	}
	logger.Debug("カレンダーから%d件のイベントを取得", len(events.Items))
	return events, nil
}
