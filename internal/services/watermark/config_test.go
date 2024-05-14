package watermark

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name       string
		expectPath string
		expectSize int
	}{
		{"new config successfully", "./web/static/root/watermark.png", 96},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Config()
			assert.Equal(t, got.Path, tt.expectPath, "got %q, want %q", got.Path, tt.expectPath)
			assert.Equal(t, got.Size, tt.expectSize, "got %q, want %q", got.Size, tt.expectSize)
		})
	}
}

func TestGetPathFromENV(t *testing.T) {
	tests := []struct {
		name   string
		want   string
		expect string
	}{
		{"set watermark path successfully", "test", "test"},
		{"no set watermark path", "", "./web/static/root/watermark.png"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("WATERMARK_PATH", tt.want)
			defer os.Unsetenv("WATERMARK_PATH")

			got := getPathFromENV()
			if got != tt.expect {
				t.Errorf("got %q, want %q", got, tt.expect)
			}
		})
	}
}

func TestGetSizeFromENV(t *testing.T) {
	testcases := []struct {
		name   string
		want   string
		expect int
	}{
		{"set watermark size successfully", "320", 320},
		{"no set watermark size", "", 96},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("WATERMARK_SIZE", tt.want)
			defer os.Unsetenv("WATERMARK_SIZE")

			got := getSizeFromENV()

			assert.Equal(t, tt.expect, got, "got %q, want %q", got, tt.expect)
		})
	}
}
