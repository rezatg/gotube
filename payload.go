package gotube

// VideoData holds detailed video information.
type VideoData struct {
	CompactVideoRenderer       *CompactVideoRenderer       `json:"compactVideoRenderer,omitempty"`
	VideoPrimaryInfoRenderer   *VideoPrimaryInfoRenderer   `json:"videoPrimaryInfoRenderer,omitempty"`
	VideoSecondaryInfoRenderer *VideoSecondaryInfoRenderer `json:"videoSecondaryInfoRenderer,omitempty"`
}

// VideoPrimaryInfoRenderer holds primary information about a video.
type VideoPrimaryInfoRenderer struct {
	Title            Title        `json:"title"`
	ViewCount        ViewCount    `json:"viewCount"`
	VideoActions     VideoActions `json:"videoActions"`
	DateText         SimpleText   `json:"dateText"`
	RelativeDateText SimpleText   `json:"relativeDateText"`
}

// VideoSecondaryInfoRenderer holds secondary information about a video.
type VideoSecondaryInfoRenderer struct {
	Owner                 Owner       `json:"owner"`
	AttributedDescription Description `json:"attributedDescription"`
}

// CompactVideoRenderer contains summary information for a video.
type CompactVideoRenderer struct {
	ID string `json:"videoId"`

	Thumbnail         Thumbnail         `json:"thumbnail"`
	Title             Title             `json:"title"`
	LongBylineText    LongBylineText    `json:"longBylineText"`
	PublishedTimeText PublishedTimeText `json:"publishedTimeText"`
	LengthText        LengthText        `json:"lengthText"`
	ViewCountText     ViewCountText     `json:"viewCountText"`

	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
}

// ViewCount represents the view count of a video.
type ViewCount struct {
	VideoViewCountRenderer struct {
		ViewCount      SimpleText `json:"viewCount"`
		ShortViewCount SimpleText `json:"shortViewCount"`
	} `json:"videoViewCountRenderer"`
}

// SimpleText represents a simple text structure.
type SimpleText struct {
	SimpleText string `json:"simpleText"`
}

// Description represents a description structure.
type Description struct {
	Content string `json:"content"`
}

// Thumbnail represents thumbnail information for a video. Contains an array of thumbnail URLs with dimensions.
type Thumbnails struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Thumbnail represents thumbnail information for a video. Contains an array of thumbnail URLs with dimensions.
type Thumbnail struct {
	Thumbnails []Thumbnails `json:"thumbnails"`
}

// Owner represents the owner of a video.
type Owner struct {
	VideoOwnerRenderer struct {
		Title struct {
			Runs [1]struct {
				Text string `json:"text"`
			} `json:"runs"`
		} `json:"title"`
		SubscriberCountText SimpleText `json:"subscriberCountText"`
	} `json:"videoOwnerRenderer"`
}

// VideoActions represents the actions available for a video.
type VideoActions struct {
	MenuRenderer struct {
		TopLevelButtons []struct {
			SegmentedLikeDislikeButtonViewModel struct {
				LikeButtonViewModel struct {
					LikeButtonViewModel struct {
						ToggleButtonViewModel struct {
							ToggleButtonViewModel struct {
								DefaultButtonViewModel struct {
									ButtonViewModel struct {
										Title string `json:"title"`
									} `json:"buttonViewModel"`
								} `json:"defaultButtonViewModel"`
							} `json:"toggleButtonViewModel"`
						} `json:"toggleButtonViewModel"`
					} `json:"likeButtonViewModel"`
				} `json:"likeButtonViewModel"`
			} `json:"segmentedLikeDislikeButtonViewModel"`
		} `json:"topLevelButtons"`
	} `json:"menuRenderer"`
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
	Runs []struct {
		Text string `json:"text"`
	} `json:"runs"`
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

// GetTitle returns the title of the video.
func (vd *VideoData) GetTitle() string {
	if title := vd.VideoPrimaryInfoRenderer.Title.Runs; len(title) > 0 {
		return title[0].Text
	}

	return vd.CompactVideoRenderer.GetTitle()
}

// GetTitle returns the title of the video.
func (cr *CompactVideoRenderer) GetTitle() string {
	if len(cr.Title.Runs) > 0 {
		return cr.Title.Runs[0].Text
	}
	return ""
}

// GetChannel returns the name of the channel.
func (vd *VideoData) GetChannel() string {
	if channel := vd.VideoSecondaryInfoRenderer.Owner.VideoOwnerRenderer.Title.Runs; len(channel) > 0 {
		return channel[0].Text
	}

	return ""
}

// GetChannel returns the name of the channel.
func (cr *CompactVideoRenderer) GetChannel() string {
	if len(cr.LongBylineText.Runs) > 0 {
		return cr.LongBylineText.Runs[0].Text
	}
	return ""
}

// GetViews extracts the video view count.
func (vd *VideoData) GetViews() string {
	return vd.VideoPrimaryInfoRenderer.ViewCount.VideoViewCountRenderer.ViewCount.SimpleText
}

// GetShortView extracts the video short view count.
func (vd *VideoData) GetShortView() string {
	return vd.VideoPrimaryInfoRenderer.ViewCount.VideoViewCountRenderer.ShortViewCount.SimpleText
}

// GetViews extracts the video view count.
func (cr *CompactVideoRenderer) GetViews() string {
	return cr.ViewCountText.SimpleText
}

// GetDuration extracts the video duration.
func (vd *VideoData) GetDuration() string {
	return vd.CompactVideoRenderer.LengthText.SimpleText
}

// GetDuration extracts the video duration.
func (cr *CompactVideoRenderer) GetDuration() string {
	return cr.LengthText.SimpleText
}

// GetPublishTime extracts the video publish time.
func (vd *VideoData) GetPublishTime() string {
	return vd.VideoPrimaryInfoRenderer.DateText.SimpleText
}

// GetRelativeDate extracts the video relative date.
func (vd *VideoData) GetRelativeDate() string {
	return vd.VideoPrimaryInfoRenderer.RelativeDateText.SimpleText
}

// GetPublishTime extracts the video publish time.
func (cr *CompactVideoRenderer) GetPublishTime() string {
	return cr.PublishedTimeText.SimpleText
}

// GetUrlSuffix extracts the video URL suffix.
func (vd *VideoData) GetUrlSuffix() string {
	return vd.CompactVideoRenderer.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

// GetUrlSuffix extracts the video URL suffix.
func (cr *CompactVideoRenderer) GetUrlSuffix() string {
	return cr.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

// GetUrl constructs the full URL for the video.
func (vd *VideoData) GetUrl() string {
	return baseURL + vd.CompactVideoRenderer.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

// GetUrl constructs the full URL for the video.
func (cr *CompactVideoRenderer) GetUrl() string {
	return baseURL + cr.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

// GetThumbnail retrieves the first thumbnail URL from the list of thumbnails associated with the video.
func (vd *VideoData) GetUrlThumbnail() string {
	if thumbnail := vd.CompactVideoRenderer.Thumbnail.Thumbnails; len(thumbnail) > 0 {
		return thumbnail[0].URL
	}

	return ""
}

// GetThumbnail retrieves the first thumbnail URL from the list of thumbnails associated with the video.
func (cr *CompactVideoRenderer) GetUrlThumbnail() string {
	if thumbnail := cr.Thumbnail.Thumbnails; len(thumbnail) > 0 {
		return thumbnail[0].URL
	}

	return ""
}

// GetThumbnails retrieves the list of all thumbnail information for the video.
func (cr *CompactVideoRenderer) GetThumbnails() []Thumbnails {
	return cr.Thumbnail.Thumbnails
}

// GetDescription extracts the video description.
func (vd *VideoData) GetDescription() string {
	return vd.VideoSecondaryInfoRenderer.AttributedDescription.Content
}

// GetLikeCount extracts the video like count.
func (vd *VideoData) GetLikeCount() string {
	if TopLevelButtons := vd.VideoPrimaryInfoRenderer.VideoActions.MenuRenderer.TopLevelButtons; len(TopLevelButtons) > 0 {
		return TopLevelButtons[0].SegmentedLikeDislikeButtonViewModel.LikeButtonViewModel.LikeButtonViewModel.ToggleButtonViewModel.ToggleButtonViewModel.DefaultButtonViewModel.ButtonViewModel.Title
	}

	return ""
}

// GetDislikeCount extracts the video dislike count.
func (vd *VideoData) GetSubscriberCount() string {
	return vd.VideoSecondaryInfoRenderer.Owner.VideoOwnerRenderer.SubscriberCountText.SimpleText
}

// GetID extracts the video ID.
func (vd *VideoData) ID() string {
	return vd.CompactVideoRenderer.ID
}
