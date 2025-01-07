package gotube

// VideoData represents the complete data for a single video.
type VideoData struct {
	ID string `json:"videoId"`

	Thumbnail         Thumbnail         `json:"thumbnail"`
	Title             Title             `json:"title"`
	LongBylineText    LongBylineText    `json:"longBylineText"`
	PublishedTimeText PublishedTimeText `json:"publishedTimeText"`
	LengthText        LengthText        `json:"lengthText"`
	ViewCountText     ViewCountText     `json:"viewCountText"`

	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
}

// Thumbnail represents thumbnail information for a video. Contains an array of thumbnail URLs with dimensions.
type Thumbnails struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Thumbnail struct {
	Thumbnails []Thumbnails `json:"thumbnails"`
}

// Title represents the title of a video, including accessibility information.
type Title struct {
	Runs []struct {
		Text string `json:"text"`
	} `json:"runs"`
	Accessibility struct {
		AccessibilityData struct {
			Label string `json:"label"`
		} `json:"accessibilityData"`
	} `json:"accessibility"`
}

// LongBylineText represents the long channel name or description, including navigation details.
type LongBylineText struct {
	Runes []struct {
		Text string `json:"text"`
	} `json:"runs"`

	NavigationEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`

		// commandMetadata, browseEndpoint
	} `json:"navigationEndpoint"`
}

// PublishedTimeText represents the simple text for the video's publish date.
type PublishedTimeText struct {
	SimpleText string `json:"simpleText"`
}

// LengthText represents the video length, including accessibility information.
type LengthText struct {
	Accessibility struct {
		AccessibilityData struct {
			Label string `json:"label"`
		} `json:"accessibilityData"`
	} `json:"accessibility"`
	SimpleText string `json:"simpleText"`
}

// ViewCountText represents the simple text for the video's view count.
type ViewCountText struct {
	SimpleText string `json:"simpleText"`
}

// NavigationEndpoint represents the navigation endpoint for the video.
type NavigationEndpoint struct {
	CommandMetadata struct {
		WebCommandMetadata struct {
			Url string `json:"url"`
		} `json:"webCommandMetadata"`
	} `json:"commandMetadata"`
}

// GetTitle extracts the video title.
func (vd *VideoData) GetTitle() string {
	return vd.Title.Runs[0].Text
}

// GetChannel extracts the channel name.
func (vd *VideoData) GetChannel() string {
	return vd.LongBylineText.Runes[0].Text
}

// GetViews extracts the video view count.
func (vd *VideoData) GetViews() string {
	return vd.ViewCountText.SimpleText
}

// GetDuration extracts the video duration.
func (vd *VideoData) GetDuration() string {
	return vd.LengthText.SimpleText
}

// GetPublishTime extracts the video publish time.
func (vd *VideoData) GetPublishTime() string {
	return vd.PublishedTimeText.SimpleText
}

// GetUrlSuffix extracts the video URL suffix.
func (vd *VideoData) GetUrlSuffix() string {
	return vd.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

// Url constructs and returns the complete URL for the video using the base URL and the video URL suffix.
func (vd *VideoData) GetUrl() string {
	return baseURL + vd.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

// GetThumbnail retrieves the first thumbnail URL from the list of thumbnails associated with the video.
func (vd *VideoData) GetThumbnailUrl() string {
	return vd.Thumbnail.Thumbnails[0].URL
}

// GetThumbnails retrieves the list of all thumbnail information for the video.
func (vd *VideoData) GetThumbnails() []Thumbnails {
	return vd.Thumbnail.Thumbnails
}
