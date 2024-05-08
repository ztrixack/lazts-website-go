package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func UpdateImagePaths(markdownData []byte, newPath string) []byte {
	// Define a regular expression to match image tags
	imageTagRegex := regexp.MustCompile(`!\[.*\]\((.*?)\)`)

	// Replace image paths using regular expressions
	newMarkdown := imageTagRegex.ReplaceAllFunc(markdownData, func(match []byte) []byte {
		// Extract the image path from the match
		path := string(imageTagRegex.FindSubmatch(match)[1])
		// Update the path if it's not an absolute path
		if !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
			updatedPath := filepath.Join(newPath, path)
			return []byte(fmt.Sprintf("![%s](%s)", path, updatedPath))
		}
		// Otherwise, return the original match
		return match
	})

	return newMarkdown
}

func UpdateFeaturedImagePaths(path string, newPath string) string {
	if !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		return filepath.Join(newPath, path)
	}

	return path
}
