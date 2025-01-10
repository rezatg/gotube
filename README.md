# gotube

`gotube` is a Go library for performing YouTube searches and extracting video data seamlessly. This library provides a simple interface to search for videos on YouTube, retrieve associated data, and facilitate further processing.

## Features

- **YouTube Search**: Easily perform searches for videos on YouTube based on specified search terms.
- **Video Data Extraction**: Extract relevant information about each video, such as video ID, title, views, duration, and more.
- **Optimized Performance**: Utilizes the [fasthttp](https://github.com/valyala/fasthttp) library for efficient HTTP requests and responses.
- **Efficient JSON Parsing**: Uses the [go-json](https://github.com/goccy/go-json) library for fast and efficient JSON parsing.


## Installation

To use the `gotube` library in your Go project, you can fetch it using `go get`:

```bash
go get github.com/rezatg/gotube
```

## Structs and Methods 
### SearchOptions Struct The SearchOptions struct defines the parameters for a YouTube search query: 
   - **`SearchTerms`**: The search query string. 
   - **`Limit`**: The maximum number of results to retrieve (not actively used in the current implementation).


### Example Usage
```go
package main

import (
   "fmt"
   "github.com/rezatg/gotube"
)

func main() {
   // Create a new GoTube instance.
   youtube := gotube.NewGoTube()

   // Perform a YouTube search with the search term "Ali Sorena, Negar" and limit the max-results to 1
   results, _ := gotube.Search(&gotube.SearchOptions{
      SearchTerms: "Ali Sorena, Negar",
      Limit: 1,
   })

   // Print details for the first result.
   videoInfo := result[0]
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
```

## Get Video Information
You can retrieve detailed information about a specific video using the GetInfoVideo method. Here is an example based on [examples/main.go](https://github.com/rezatg/gotube/blob/master/examples/getInfo/main.go):

```go
package main

import (
   "fmt"
   "log"

   "github.com/rezatg/gotube"
)

func main() {
   // Create a new GoTube instance.
   youtube := gotube.NewGoTube()

   // Specify the YouTube video URL.  
   videoURL := "https://youtu.be/1r8sEJTtwzE?si=jOvHjVawN-2cgQKi"


   // Retrieve video information.
   videoInfo, err := youtube.GetInfoVideo(videoURL)
   if err != nil {
      log.Fatalf("Error getting video info: %v", err)
   }

   // Access and print the desired information.  See below for a full list of available fields.
   fmt.Println("Video Title:", videoInfo.GetTitle())
   fmt.Println("Channel:", videoInfo.GetChannel())
   fmt.Println("Views:", videoInfo.GetViews())
   fmt.Println("Description:", videoInfo.GetDescription())
   // ... access other fields as needed
}
```

## Available Fields **`(GetInfoVideo)`**:
The **GetInfoVideo** method returns a VideoData struct, which provides access to a variety of video details through its methods :

  - `ID()`: Returns the video ID.
  - `GetTitle()`: Returns the video title.
  - `GetDescription()`: Returns the video description.
  - `GetChannel()`: Returns the channel name.
  - `GetUrlThumbnail()`: Returns the URL of the video thumbnail.
  - `GetDuration()`: Returns the video duration.
  - `GetLikeCount()`: Returns the like count.
  - `GetPublishTime()`: Returns the publish time.
  - `GetRelativeDate()`: Returns the relative publish date (e.g., "2 weeks ago").
  - `GetViews()`: Returns the view count.
  - `GetShortView()`: Returns the abbreviated view count (e.g., "1.2M views").
  - `GetSubscriberCount()`: Returns the subscriber count of the channel.
  - `GetUrlSuffix()`: Returns the URL suffix of the video.
  - `GetUrl()`: Returns the full URL of the video.


## Error Handling
The GetInfoVideo method may return errors in the following cases:
   - **Invalid Video URL**: If the provided URL is incorrect or malformed.
   - **Network Issues**: If there are connectivity issues or server errors.  

> It is recommended to handle errors appropriately, as shown in the examples.
