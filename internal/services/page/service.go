package page

import (
	"html/template"
	"io"
	"lazts/internal/modules/log"
	"lazts/internal/modules/markdown"
)

type Servicer interface {
	Render(wr io.Writer, page string) error
	RenderMarkdown(wr io.Writer, domain string, page string) error
	RenderBlackhole(wr io.Writer, counter int) error
	RenderCloud(wr io.Writer, counter int) error
}

type service struct {
	log       log.Moduler
	markdown  markdown.Moduler
	templates *template.Template
}

var _ Servicer = (*service)(nil)

const HTML_PATH = "web/templates/**/*.html"

func New(log log.Moduler, md markdown.Moduler) *service {
	tpl, err := template.ParseGlob(HTML_PATH)
	if err != nil {
		log.Err(err).C("Error setting up templates")
	}

	return &service{
		log:       log,
		markdown:  md,
		templates: tpl,
	}
}
