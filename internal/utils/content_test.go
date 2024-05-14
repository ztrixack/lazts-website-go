package utils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContentDir(t *testing.T) {
	assert.Equal(t, filepath.Join(".", "contents"), GetContentDir(), "Should handle no additional directories correctly")
	assert.Equal(t, filepath.Join(".", "contents", "images"), GetContentDir("images"), "Should correctly append directory names")
	assert.Equal(t, filepath.Join(".", "contents", "images", "2021"), GetContentDir("images", "2021"), "Should correctly append multiple directory names")
}
