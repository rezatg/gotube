package gotube

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/isaacpd/costanza/pkg/util"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// baseURL defines the base URL for YouTube searches.
var baseURL string = "https://youtube.com"

// optionsSearch represents the options for a YouTube search query.
// SearchTerms: The search query string.
// Limit: The maximum number of results to retrieve (currently not used in this implementation).
type OptionsSearch struct {
	SearchTerms string
	Limit       int
}

// Search performs a YouTube search using the provided options.
// It constructs the search URL, sends an HTTP GET request, and prints the raw response body.
// Note: This function currently doesn't handle pagination or result parsing, it only retrieves the raw HTML.
func Search(opt *OptionsSearch) ([]VideoData, error) {
	encodedSearch := url.QueryEscape(opt.SearchTerms)                        // Encode the search terms for use in the URL.
	url := fmt.Sprintf("%s/results?search_query=%s", baseURL, encodedSearch) // Construct the YouTube search URL.

	req := fasthttp.AcquireRequest()   // Acquire a request object from the fasthttp pool. This helps to improve performance by reusing objects.
	defer fasthttp.ReleaseRequest(req) // Release the request object back to the pool when finished. This is crucial for resource management.

	req.SetRequestURI(url) // Set the request URI and method.
	req.Header.SetMethod("GET")
	req.Header.Add("Accept-Language", "en") // Add an Accept-Language header for better compatibility.

	resp := fasthttp.AcquireResponse()   // Acquire a response object from the fasthttp pool. Similar to the request, this improves performance.
	defer fasthttp.ReleaseResponse(resp) // Release the response object back to the pool when finished. Essential for efficient memory usage.

	// Execute the HTTP request, handling redirects. The util.DoWithRedirects function is assumed to be defined elsewhere and handles redirects efficiently.
	if err := util.DoWithRedirects(req, resp); err != nil {
		fmt.Printf("Error getting youtube search results: %s", err)
		return nil, err
	}

	results, err := ParseHtml(string(resp.Body()), opt.Limit)
	// results[0].get
	return results, err
}

// parseHtml parses the HTML response to extract video data.
func ParseHtml(response string, limit int) ([]VideoData, error) {
	// Find and extract ytInitialData JSON
	var startIndex int = strings.Index(response, "ytInitialData")
	if startIndex == -1 {
		return nil, fmt.Errorf("ytInitialData not found")
	}
	startIndex += len("ytInitialData") + 3 // +3 for " = {"
	endIndex := strings.Index(response[startIndex:], "};") + startIndex + 1
	if endIndex == -1 {
		return nil, fmt.Errorf("end of ytInitialData not found")
	}

	var ytInitialData struct {
		Contents json.RawMessage `json:"contents"` // Holds the contents of ytInitialData
	}

	// Unmarshal the JSON data into ytInitialData structure
	if err := json.Unmarshal([]byte(response[startIndex:endIndex]), &ytInitialData); err != nil {
		return nil, fmt.Errorf("error unmarshalling ytInitialData: %w", err)
	}

	var p fastjson.Parser
	v, err := p.ParseBytes(ytInitialData.Contents) // Parse the contents using fastjson
	if err != nil {
		return nil, err
	}

	var results = []VideoData{}
	for _, contents := range v.GetArray("twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer", "contents") {
		for _, video := range contents.GetArray("itemSectionRenderer", "contents") {
			if limit >= 1 && len(results) >= limit {
				return results, nil
			}

			// Check if videoRenderer exists
			if video.Exists("videoRenderer") {
				var videoData VideoData
				// Unmarshal the videoRenderer object into VideoData structure
				if err = json.Unmarshal([]byte(video.GetObject("videoRenderer").String()), &videoData); err != nil {
					return nil, err
				}
				// Append the extracted VideoData to results
				results = append(results, videoData)
			}
		}
	}

	return results, nil
}
