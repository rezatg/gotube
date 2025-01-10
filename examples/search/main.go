package main

import (
	"fmt"

	"github.com/rezatg/gotube"
)

func main() {
	// Create a new instance of GoTube
	youtube := gotube.NewGoTube()

	// Search for videos with the specified search terms and limit the result to 1 video
	result, err := youtube.Search(&gotube.OptionsSearch{
		SearchTerms: "Ali Sorena, Negar",
		Limit:       1,
	})
	if err != nil {
		fmt.Printf("Error occurred while searching: %e\n", err)
		return
	}

	// Print details for the first result.
	videoInfo := result[0]
	fmt.Println("Video Information:")
	fmt.Printf("ID: %s\nTitle: %s\nThumbnail URL: %s\nThumbnails: %v\nChannel: %s\nDuration: %s\nViews: %s\nURL Suffix: %s\nURL: %s\nPublish Time: %s\n",
		videoInfo.ID,
		videoInfo.GetTitle(),
		videoInfo.GetUrlThumbnail(),
		videoInfo.GetThumbnails(),
		videoInfo.GetChannel(),
		videoInfo.GetDuration(),
		videoInfo.GetViews(),
		videoInfo.GetUrlSuffix(),
		videoInfo.GetUrl(),
		videoInfo.GetPublishTime(),
	)
}
