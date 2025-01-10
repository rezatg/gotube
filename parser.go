package gotube

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/goccy/go-json"
	"github.com/valyala/fastjson"
)

// Regular expression for extracting ytInitialData JSON from the HTML response.
var ytInitialDataRegex *regexp.Regexp = regexp.MustCompile(`var ytInitialData\s*=\s*(\{.+?\});`)

// ParseHtmlSearch extracts video data from the parsed JSON response.
// It takes the HTML response and a limit for the number of videos to extract.
func ParseHtmlSearch(response []byte, limit int) ([]CompactVideoRenderer, error) {
	var resp = string(response)
	data := ytInitialDataRegex.FindStringSubmatch(resp)

	// Find and extract ytInitialData JSON
	if len(data) < 2 {
		return nil, fmt.Errorf("ytInitialData not found")
	}

	// Parse the JSON using fastjson for performance.
	var p fastjson.Parser
	v, err := p.Parse(data[1])
	if err != nil {
		return nil, err
	}

	results := make([]CompactVideoRenderer, 0, limit)
	for _, contents := range v.GetArray("contents", "twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer", "contents") {
		for _, video := range contents.GetArray("itemSectionRenderer", "contents") {
			// Check if videoRenderer exists
			videoRenderer := video.Get("videoRenderer")
			if videoRenderer == nil {
				continue // Skip if videoRenderer is not found
			}

			// Unmarshal the videoRenderer object into CompactVideoRenderer structure
			var videoData CompactVideoRenderer
			if err = json.Unmarshal([]byte(videoRenderer.String()), &videoData); err != nil {
				return nil, err
			}

			// Append the extracted CompactVideoRenderer to results
			results = append(results, videoData)
			if len(results) == limit {
				break // Exit early if limit reached
			}

		}
	}

	if len(results) == 0 {
		return nil, errors.New("no videos found")
	}

	return results, nil
}

// ParseHtmlInfoVideo extracts detailed video information from the parsed JSON response.
// It takes the HTML response and returns a VideoData struct containing the video's information.
func ParseHtmlInfoVidoe(response []byte) (*VideoData, error) {
	var (
		resp string   = string(response)
		data []string = ytInitialDataRegex.FindStringSubmatch(resp)
	)

	// Find and extract ytInitialPlayerResponse JSON
	if len(data) < 2 {
		return nil, fmt.Errorf("ytInitialData not found")
	}

	// Parse the JSON using fastjson for performance.
	var p fastjson.Parser
	v, err := p.Parse(data[1])
	if err != nil {
		return nil, err
	}

	cont := v.Get("contents", "twoColumnWatchNextResults")
	if cont == nil {
		return nil, errors.New("contents.twoColumnWatchNextResults not found")
	}

	var videoData VideoData
	for _, contents := range cont.GetArray("results", "results", "contents") {
		if contents.Exists("videoPrimaryInfoRenderer") || contents.Exists("videoSecondaryInfoRenderer") {
			if err = json.Unmarshal([]byte(contents.String()), &videoData); err != nil {
				return nil, fmt.Errorf("failed to unmarshal video info: %e", err)
			}
		}
	}

	for _, contents := range cont.GetArray("secondaryResults", "secondaryResults", "results") {
		if contents.Exists("compactVideoRenderer") {
			if err = json.Unmarshal([]byte(contents.String()), &videoData); err != nil {
				return nil, fmt.Errorf("failed to unmarshal video info: %e", err)
			}
			break
		}
	}

	return &videoData, nil
}
