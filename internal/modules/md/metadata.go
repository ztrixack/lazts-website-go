package md

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

func (m *module) Metadata(source []byte, result interface{}) error {
	context := parser.NewContext()
	markdown := goldmark.New(goldmark.WithExtensions(meta.New()))

	var buf bytes.Buffer
	if err := markdown.Convert(source, &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}

	metadata := meta.Get(context)
	metadata["ReadTime"] = calculateReadTime(source)

	dataByte, _ := json.Marshal(metadata)
	json.Unmarshal(dataByte, result)

	return nil
}

func calculateReadTime(markdown []byte) int {
	text := removeMarkdownFormatting(string(markdown))
	words := strings.Fields(text)
	wordCount := len(words)
	const wordsPerMinute = 250
	readTime := wordCount / wordsPerMinute
	if wordCount%wordsPerMinute != 0 {
		readTime++
	}

	return readTime
}

func removeMarkdownFormatting(text string) string {
	urlRegex := regexp.MustCompile(`\[(.*?)\]\(http.*?\)`)
	text = urlRegex.ReplaceAllString(text, "$1")
	symbols := []string{"#", "*", "_", "`"}
	for _, symbol := range symbols {
		text = strings.ReplaceAll(text, symbol, "")
	}

	return text
}
