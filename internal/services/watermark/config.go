package watermark

import (
	"os"
	"strconv"
)

type config struct {
	Path string
	Size int
}

func Config() *config {
	path := getPathFromENV()
	size := getSizeFromENV()

	return &config{
		Path: path,
		Size: size,
	}
}

func getPathFromENV() string {
	path := os.Getenv("WATERMARK_PATH")
	if path == "" {
		path = "./web/static/root/watermark.png"
	}
	return path
}

func getSizeFromENV() int {
	size := os.Getenv("WATERMARK_SIZE")
	if s, _ := strconv.Atoi(size); s > 0 {
		return s
	}
	return 48
}
