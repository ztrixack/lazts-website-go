package utils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContentDir(t *testing.T) {
	assert.Equal(t, filepath.Join(".", "web", "contents"), GetContentDir(), "Should handle no additional directories correctly")
	assert.Equal(t, filepath.Join(".", "web", "contents", "images"), GetContentDir("images"), "Should correctly append directory names")
	assert.Equal(t, filepath.Join(".", "web", "contents", "images", "2021"), GetContentDir("images", "2021"), "Should correctly append multiple directory names")
}

func TestGetContentPath(t *testing.T) {
	assert.Equal(t, filepath.Join("/", "static", "contents"), GetContentPath(), "Should handle no additional directories correctly")
	assert.Equal(t, filepath.Join("/", "static", "contents", "images"), GetContentPath("images"), "Should correctly append directory names")
	assert.Equal(t, filepath.Join("/", "static", "contents", "images", "2021"), GetContentPath("images", "2021"), "Should correctly append multiple directory names")
}

func TestGetStaticDir(t *testing.T) {
	assert.Equal(t, filepath.Join(".", "web", "static"), GetStaticDir(), "Should handle no additional directories correctly")
	assert.Equal(t, filepath.Join(".", "web", "static", "images"), GetStaticDir("images"), "Should correctly append directory names")
	assert.Equal(t, filepath.Join(".", "web", "static", "images", "2021"), GetStaticDir("images", "2021"), "Should correctly append multiple directory names")
}

func TestGetStaticPath(t *testing.T) {
	assert.Equal(t, filepath.Join("/", "static"), GetStaticPath(), "Should handle no additional directories correctly")
	assert.Equal(t, filepath.Join("/", "static", "images"), GetStaticPath("images"), "Should correctly append directory names")
	assert.Equal(t, filepath.Join("/", "static", "images", "2021"), GetStaticPath("images", "2021"), "Should correctly append multiple directory names")
}
