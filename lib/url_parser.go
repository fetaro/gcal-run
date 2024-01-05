package lib

import (
	"regexp"
	"strings"

	"google.golang.org/api/calendar/v3"
)

type URLParser struct {
}

func NewURLParser() *URLParser {
	return &URLParser{}
}

func (p *URLParser) Parse(event *calendar.Event) (string, string) {
	url := parseZoomUrl(event)
	if url != "" {
		return "Zoom", url
	}
	url = parseMeetUrl(event)
	if url != "" {
		return "Meet", url
	}
	url = parseTeamsUrl(event)
	if url != "" {
		return "Teams", url
	}
	return "Other", ""
}

func parseMeetUrl(event *calendar.Event) string {
	if event.ConferenceData == nil {
		return ""
	}
	for _, entryPoint := range event.ConferenceData.EntryPoints {
		if entryPoint.EntryPointType == "video" {
			return entryPoint.Uri
		}
	}
	return ""
}

func parseZoomUrl(event *calendar.Event) string {
	if event.Location != "" && strings.Contains(event.Location, "zoom.us") {
		return event.Location
	}
	tagMatcher := regexp.MustCompile(`<a href=\"(https://zoom.us/.*?)\">.*`)
	if tagMatcher.MatchString(event.Description) {
		matched := tagMatcher.FindStringSubmatch(event.Description)
		return matched[1]
	}
	return ""
}

func parseTeamsUrl(event *calendar.Event) string {
	tagMatcher := regexp.MustCompile(`<(https://teams.microsoft.com/l/meetup-join.*)>`)
	if tagMatcher.MatchString(event.Description) {
		matched := tagMatcher.FindStringSubmatch(event.Description)
		return matched[1]
	}
	return ""
}
