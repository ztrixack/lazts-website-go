package markdown

import (
	"bytes"
	"context"
	"fmt"
	"lazts/pkg/logger"
	"os"

	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	_ "github.com/yuin/goldmark-emoji/definition"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
	"go.abhg.dev/goldmark/mermaid/mermaidcdp"
	"go.abhg.dev/goldmark/toc"

	katex "github.com/FurqanSoftware/goldmark-katex"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"

	_ "embed"
)

//go:embed mermaid.min.js
var mermaidJSSource string

type Servicer interface {
	RenderMarkdownToHTML(ctx context.Context, name string) ([]byte, error)
}

type service struct {
	log logger.Logger
}

func New(log logger.Logger) *service {
	return &service{
		log,
	}
}

func (s *service) RenderMarkdownToHTML(ctx context.Context, name string) ([]byte, error) {
	filepath := fmt.Sprintf("./static/%s.md", name)
	markdownData, err := os.ReadFile(filepath)
	if err != nil {
		s.log.Err(err).Fields("filepath", filepath).E("failed to read markdown file")
		return nil, err
	}

	compiler, err := mermaidcdp.New(&mermaidcdp.Config{JSSource: mermaidJSSource})
	if err != nil {
		s.log.Err(err).E("failed to create mermaid compiler")
		return nil, err
	}
	defer compiler.Close()

	context := parser.NewContext()

	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, emoji.Emoji, figure.Figure,
			meta.New(meta.WithTable()),
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			&katex.Extender{}, &toc.Extender{},
			&mermaid.Extender{Compiler: compiler}),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithHardWraps(), html.WithXHTML(), html.WithUnsafe()),
	)

	buffer := new(bytes.Buffer)
	if err := markdown.Convert(markdownData, buffer, parser.WithContext(context)); err != nil {
		s.log.Err(err).E("failed to convert markdown to html")
		return nil, err
	}

	return buffer.Bytes(), nil
}
