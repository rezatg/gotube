package gotube_test

import (
	"os"
	"testing"

	"github.com/rezatg/gotube"
	"github.com/stretchr/testify/assert"
)

func readMockJsonFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func TestParseHtml(t *testing.T) {
	// Test case 1: Valid JSON data
	jsonData := readMockJsonFile("result.txt")
	results, err := gotube.ParseHtml(jsonData, 1)

	assert.NoError(t, err)
	assert.Len(t, results, 1)

	assert.Equal(t, "testVideo1", results[0].ID)
}
