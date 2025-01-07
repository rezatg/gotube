# gotube

`gotube` is a Go library for performing YouTube searches and extracting video data seamlessly. This library provides a simple interface to search for videos on YouTube, retrieve associated data, and facilitate further processing.

## Features

- **YouTube Search**: Easily perform searches for videos on YouTube based on specified search terms.
- **Video Data Extraction**: Extract relevant information about each video, such as video ID, title, views, duration, and more.
- **Optimized Performance**: Utilizes the `fasthttp` library for efficient HTTP requests and responses.

## Installation

To use the `gotube` library in your Go project, you can fetch it using `go get`:

```bash
go get github.com/rezatg/gotube
```

### OptionsSearch Struct

The OptionsSearch struct defines the options for a YouTube search query, including:
- SearchTerms: The search query string.
- Limit: The maximum number of results to retrieve (not actively used in the current implementation).

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/rezatg/gotube"
)

func main() {
    // Perform a YouTube search with the search term "Ali Sorena, Negar" and limit the max-results to 1
    results, _ := gotube.Search(&gotube.OptionsSearch{
        SearchTerms: "Ali Sorena, Negar",
        Limit: 1,
    })

    // Output specific information about the first video from the search results
    fmt.Println(results[0].ID)
	fmt.Println(results[0].GetTitle())
	fmt.Println(results[0].GetThumbnailUrl())
	fmt.Println(results[0].GetThumbnails())
	fmt.Println(results[0].GetChannel())
	fmt.Println(results[0].GetDuration())
	fmt.Println(results[0].GetViews())
	fmt.Println(results[0].GetUrlSuffix())
	fmt.Println(results[0].GetUrl())
    fmt.Println(results[0].GetPublishTime())
}
```