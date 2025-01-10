package gotube_test

import (
	"os"
	"testing"

	"github.com/rezatg/gotube"
)

// import (
// 	"os"
// 	// "github.com/rezatg/gotube"
// )

func readMockJsonFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func TestParseHtml(t *testing.T) {
	jsonData := readMockJsonFile("./test/response.txt")
	results, err := gotube.ParseHtmlSearch(jsonData, 1)
	println(111, results, err.Error())
}

// 	// Test case 1: Valid JSON data
// 	jsonData := readMockJsonFile("./temp/result.txt")
// 	results, err := gotube.ParseHtmlSearch(jsonData, 1)

// 	assert.NoError(t, err)
// 	assert.Len(t, results, 1)

// 	assert.Equal(t, "testVideo1", results[0].ID)
// }
