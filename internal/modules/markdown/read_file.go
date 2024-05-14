package markdown

import (
	"os"

	_ "github.com/yuin/goldmark-emoji/definition"

	"lazts/internal/utils"
)

func (m *module) ReadFile(domain string, name string) ([]byte, error) {
	file := utils.GetContentDir(domain, name, "page.md")
	markdownData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return utils.UpdateImagePaths(markdownData, utils.GetContentPath(domain, name)), nil
}
