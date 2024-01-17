package lib

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
	urlPrefixList := []string{"https://zoom.us", "https://teams.microsoft.com/l/meetup-join", "https://meet.google.com"}
	for _, urlPrefix := range urlPrefixList {
		if strings.Contains(event.Location, urlPrefix) {
			return event.Location, nil
		}
		// DescriptionにURLが含まれている場合は採用
		descriptionRegExpList := []string{
			"<a href=\"(" + urlPrefix + ".*?)\">.*",
			"(" + urlPrefix + `[\.0-9a-zA-Z-_?/=%&@]+` + ")",
		}
		for _, descriptionRegExp := range descriptionRegExpList {
			tagMatcher := regexp.MustCompile(descriptionRegExp)
			if tagMatcher.MatchString(event.Description) {
				matched := tagMatcher.FindStringSubmatch(event.Description)
				return matched[1], nil
			}
		}
	}
	return "", fmt.Errorf("event is not online meeting")
}
