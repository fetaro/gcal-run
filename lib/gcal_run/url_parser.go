package gcal_run

import (
	"fmt"
	"regexp"
	"strings"

	"google.golang.org/api/calendar/v3"
)

type URLParser struct {
}

func NewURLParser() *URLParser {
	return &URLParser{}
}

func (p *URLParser) Parse(event *calendar.Event) (string, error) {
	// EntryPointTypeがvideoなら採用。meetはこれでマッチする
	if event.ConferenceData != nil && event.ConferenceData.EntryPoints != nil {
		for _, entryPoint := range event.ConferenceData.EntryPoints {
			if entryPoint.EntryPointType == "video" {
				return entryPoint.Uri, nil
			}
		}
	}
	// LocationにURLが含まれている場合は採用
	locationUrlPatternList := []string{
		"https://zoom.us",
		"https://us02web.zoom.us",
		"https://teams.microsoft.com/l/meetup-join",
		"https://meet.google.com",
	}
	for _, urlPrefix := range locationUrlPatternList {
		if strings.Contains(event.Location, urlPrefix) {
			return event.Location, nil
		}
	}
	// DescriptionにURLが含まれている場合は採用
	descriptionUrlPatternList := []string{
		"<(https://zoom.us.*)>",
		"<(https://teams.microsoft.com/l/meetup-join.*)>",
		"(https://meet.google.com/[a-z-]+)",
		"(http://meet.google.com/[a-z-]+)",
	}
	for _, urlRegExp := range descriptionUrlPatternList {
		tagMatcher := regexp.MustCompile(urlRegExp)
		if tagMatcher.MatchString(event.Description) {
			matched := tagMatcher.FindStringSubmatch(event.Description)
			return matched[1], nil
		}
	}
	// DescriptionにURLが含まれている場合は採用
	return "", fmt.Errorf("LocationもしくはDescriptionにTV会議のURLが含まれていません.Location=%s, Description=%s", event.Location, event.Description)
}
