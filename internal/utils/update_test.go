package utils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateImagePaths(t *testing.T) {
	testcases := []struct {
		name     string
		input    []byte
		newPath  string
		expected string
	}{
		{
			name:     "Normal Image Paths",
			input:    []byte(`This is a test image: ![test image](images/test.png) and ![another image](/images/2021/img.jpg)`),
			newPath:  "/new/path",
			expected: `This is a test image: ![images/test.png](/new/path/images/test.png) and ![another image](/images/2021/img.jpg)`,
		},
		{
			name:     "With Absolute URLs",
			input:    []byte(`This is a test image: ![test image](http://example.com/images/test.png)`),
			newPath:  "/new/path",
			expected: `This is a test image: ![test image](http://example.com/images/test.png)`,
		},
		{
			name:     "No Image Tags",
			input:    []byte(`This is text with no images.`),
			newPath:  "/new/path",
			expected: `This is text with no images.`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := UpdateImagePaths(tc.input, tc.newPath)
			assert.Equal(t, tc.expected, string(result), "Output should match expected")
		})
	}
}

func TestUpdateFeaturedImagePaths(t *testing.T) {
	testcases := []struct {
		name     string
		path     string
		newPath  string
		expected string
	}{
		{
			name:     "Relative Path",
			path:     "images/test.png",
			newPath:  "/new/path",
			expected: filepath.Join("/new/path", "images/test.png"),
		},
		{
			name:     "With Absolute URL",
			path:     "http://example.com/images/test.png",
			newPath:  "/new/path",
			expected: "http://example.com/images/test.png",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := UpdateFeaturedImagePaths(tc.path, tc.newPath)
			assert.Equal(t, tc.expected, result, "Output should match expected")
		})
	}
}
