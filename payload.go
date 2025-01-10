package gotube

// VideoData holds detailed video information.
type VideoData struct {
	CompactVideoRenderer       CompactVideoRenderer       `json:"compactVideoRenderer"`
	VideoPrimaryInfoRenderer   VideoPrimaryInfoRenderer   `json:"videoPrimaryInfoRenderer"`
	VideoSecondaryInfoRenderer VideoSecondaryInfoRenderer `json:"videoSecondaryInfoRenderer"`
}

type VideoPrimaryInfoRenderer struct {
	Title        Title        `json:"title"`
	ViewCount    ViewCount    `json:"viewCount"`
	VideoActions VideoActions `json:"videoActions"`
	DateText     struct {
		SimpleText string `json:"simpleText"`
	} `json:"dateText"`
	RelativeDateText struct {
		SimpleText string `json:"simpleText"`
	} `json:"relativeDateText"`
}

type VideoSecondaryInfoRenderer struct {
	Owner                 Owner `json:"owner"`
	AttributedDescription struct {
		Content string `json:"content"`
	} `json:"attributedDescription"`
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

type ViewCount struct {
	VideoViewCountRenderer struct {
		ViewCount struct {
			SimpleText string `json:"simpleText"`
		} `json:"viewCount"`
		ShortViewCount struct {
			SimpleText string `json:"simpleText"`
		} `json:"shortViewCount"`
	} `json:"videoViewCountRenderer"`
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

type Owner struct {
	VideoOwnerRenderer struct {
		Title struct {
			Runs [1]struct {
				Text string `json:"text"`
			} `json:"runs"`
		} `json:"title"`
		SubscriberCountText struct {
			SimpleText string `json:"simpleText"`
		} `json:"subscriberCountText"`
	} `json:"videoOwnerRenderer"`
}

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
										Title string
									}
								}
							}
						}
					}
				}
			}
		}
	}
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

// title := vd.VideoPrimaryInfoRenderer.VideoActions.MenuRenderer.TopLevelButtons
// if len(title) > 0 {
// 	return title[0].SegmentedLikeDislikeButtonViewModel.LikeButtonViewModel.LikeButtonViewModel.ToggleButtonViewModel.ToggleButtonViewModel.DefaultButtonViewModel.ButtonViewModel.Title
// }

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

func (vd *VideoData) GetShortView() string {
	return vd.VideoPrimaryInfoRenderer.ViewCount.VideoViewCountRenderer.ShortViewCount.SimpleText
}

func (cr *CompactVideoRenderer) GetViews() string {
	return cr.ViewCountText.SimpleText
}

// GetDuration extracts the video duration.
func (vd *VideoData) GetDuration() string {
	return vd.CompactVideoRenderer.LengthText.SimpleText
}

func (cr *CompactVideoRenderer) GetDuration() string {
	return cr.LengthText.SimpleText
}

// GetPublishTime extracts the video publish time.
func (vd *VideoData) GetPublishTime() string {
	return vd.VideoPrimaryInfoRenderer.DateText.SimpleText
}

func (vd *VideoData) GetRelativeDate() string {
	return vd.VideoPrimaryInfoRenderer.RelativeDateText.SimpleText
}

func (cr *CompactVideoRenderer) GetPublishTime() string {
	return cr.PublishedTimeText.SimpleText
}

// GetUrlSuffix extracts the video URL suffix.
func (vd *VideoData) GetUrlSuffix() string {
	return vd.CompactVideoRenderer.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

func (cr *CompactVideoRenderer) GetUrlSuffix() string {
	return cr.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

// GetUrl constructs the full URL for the video.
func (vd *VideoData) GetUrl() string {
	return baseURL + vd.CompactVideoRenderer.NavigationEndpoint.CommandMetadata.WebCommandMetadata.Url
}

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

func (vd *VideoData) GetDescription() string {
	return vd.VideoSecondaryInfoRenderer.AttributedDescription.Content
}

func (vd *VideoData) GetLikeCount() string {
	if TopLevelButtons := vd.VideoPrimaryInfoRenderer.VideoActions.MenuRenderer.TopLevelButtons; len(TopLevelButtons) > 0 {
		return TopLevelButtons[0].SegmentedLikeDislikeButtonViewModel.LikeButtonViewModel.LikeButtonViewModel.ToggleButtonViewModel.ToggleButtonViewModel.DefaultButtonViewModel.ButtonViewModel.Title
	}

	return ""
}

func (vd *VideoData) GetSubscriberCount() string {
	return vd.VideoSecondaryInfoRenderer.Owner.VideoOwnerRenderer.SubscriberCountText.SimpleText
}

func (vd *VideoData) ID() string {
	return vd.CompactVideoRenderer.ID
}
