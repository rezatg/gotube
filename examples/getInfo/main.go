package main

import (
	"fmt"

	"github.com/rezatg/gotube"
)

// URL of the video for which information is to be retrieved
var url string = "https://youtu.be/1r8sEJTtwzE?si=jOvHjVawN-2cgQKi"

func main() {
	// Create a new instance of GoTube
	youtube := gotube.NewGoTube()

	// Get video information using the provided URL
	result, err := youtube.GetInfoVideo(url)
	if err != nil {
		fmt.Printf("Error retrieving video information: %v\n", err)
		return
	}

	fmt.Printf("Video ID: %s\nTitle: %s\nDescription: %s\nChannel: %s\nThumbnail URL: %s\nDuration: %s\nLike Count: %s\nPublish Time: %s\nRelative Date: %s\nViews: %s\nShort View: %s\nSubscriber Count: %s\nURL Suffix: %s\nVideo URL: %s\n",
		result.ID(),
		result.GetTitle(),
		result.GetDescription(),
		result.GetChannel(),
		result.GetUrlThumbnail(),
		result.GetDuration(),
		result.GetLikeCount(),
		result.GetPublishTime(),
		result.GetRelativeDate(),
		result.GetViews(),
		result.GetShortView(),
		result.GetSubscriberCount(),
		result.GetUrlSuffix(),
		result.GetUrl(),
	)
}
