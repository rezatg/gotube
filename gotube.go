package gotube

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/valyala/fasthttp"
)

const (
	baseURL      string = "https://youtube.com" // BaseURL defines the base URL for YouTube.
	maxRedirects int    = 5                     // Maximum number of allowed redirects.
)

var (
	// Regular expression for extracting ytInitialData JSON from the HTML response.
	ytInitialDataRegex *regexp.Regexp = regexp.MustCompile(`var ytInitialData\s*=\s*(\{.+?\});`)

	// Regular expression to match YouTube video URLs. This is a relatively strict pattern. Consider refining it if needed.
	youtubeURLRegex = regexp.MustCompile(`^((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube(-nocookie)?\.com|youtu.be))(\/(?:[\w\-]+\?v=|embed\/|v\/)?)([\w\-]+)(\S+)?$`)
)

// GoTube is a wrapper structure to manage default context and timeout for requests.
type GoTube struct {
	client *fasthttp.Client // HTTP client for sending requests.
}

// NewGoTube creates a new instance of GoTube with a default timeout.
func NewGoTube() *GoTube {
	return &GoTube{
		client: &fasthttp.Client{},
	}
}

// OptionsSearch defines the options for a YouTube search query.
type OptionsSearch struct {
	SearchTerms string // Search query string
	Limit       int    // Maximum number of results to retrieve
}

// Search performs a YouTube search using the provided options.
// It constructs the search URL, sends an HTTP GET request, and prints the raw response body.
// Note: This function currently doesn't handle pagination or result parsing, it only retrieves the raw HTML.
func (gt GoTube) Search(opt *OptionsSearch) ([]CompactVideoRenderer, error) {
	if opt.SearchTerms == "" {
		return nil, fmt.Errorf("search terms cannot be empty")
	}

	// Construct the YouTube search URL.
	encodedSearch := url.QueryEscape(opt.SearchTerms)
	url := fmt.Sprintf("%s/results?search_query=%s", baseURL, encodedSearch)

	response, err := gt.sendRequest(url)
	if err != nil {
		return nil, err
	}

	return ParseHtmlSearch(response, opt.Limit)
}

// GetInfoVideo retrieves video information from a given YouTube video URL.
// It sends an HTTP GET request to the URL, parses the HTML response, and returns a VideoData struct containing the video's information.
// Note: This function relies on the HTML structure of the YouTube video page. Changes to YouTube's HTML may break this function. Error handling is implemented for network requests, but not for parsing errors within ParseHtmlInfoVideo. Consider adding more robust error handling in the future.
func (gt GoTube) GetInfoVideo(url string) (*VideoData, error) {
	// Check if the URL matches the YouTube URL pattern using a regular expression.
	if !youtubeURLRegex.MatchString(url) {
		return nil, fmt.Errorf("invalid YouTube URL: %s", url)
	}

	response, err := gt.sendRequest(url)
	if err != nil {
		return nil, err
	}

	return ParseHtmlInfoVidoe(response)
}

// sendRequest sends an HTTP GET request and handles redirects.
func (gt GoTube) sendRequest(url string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url) // Set the request URI and method.
	req.Header.SetMethod("GET")
	req.Header.Add("Accept-Language", "en")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// doWithRedirects handles HTTP redirects for a request.
	if err := gt.client.DoRedirects(req, resp, maxRedirects); err != nil {
		return nil, fmt.Errorf("error fetching URL (%s): %v", url, err)
	}

	return resp.Body(), nil
}
