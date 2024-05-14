package md

import (
	"os"
	"path/filepath"

	_ "github.com/yuin/goldmark-emoji/definition"

	"lazts/internal/utils"
)

func (m *module) ReadFile(domain string, name string) ([]byte, error) {
	file := filepath.Join(".", "contents", domain, name, "page.md")
	markdownData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return utils.UpdateImagePaths(markdownData, filepath.Join("/static", "contents", domain, name)), nil
}
