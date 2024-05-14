package markdown

import (
	_ "embed"
	"io"
)

//go:embed mermaid.min.js
var mermaidJSSource string

type Moduler interface {
	ReadFile(domain, name string) ([]byte, error)
	Convert(source []byte, wr io.Writer) error
	Metadata(source []byte, result interface{}) error
}

type module struct {
	config *config
}

var _ Moduler = (*module)(nil)

func New(config *config) *module {
	return &module{
		config: config,
	}
}

func (m *module) Config() config {
	return *m.config
}
