package vacation

import (
	"html/template"
	"io"
	"lazts/internal/modules/log"
	"lazts/internal/modules/markdown"
)

type Servicer interface {
	RenderHero(wr io.Writer) error
	RenderList(wr io.Writer) error
	RenderHighlight(wr io.Writer) error
}

type service struct {
	logger     log.Moduler
	markdowner markdown.Moduler
	templates  *template.Template
}

var _ Servicer = (*service)(nil)

const HTML_PATH = "web/templates/sections/vacations/*.html"

func New(log log.Moduler, md markdown.Moduler) *service {
	tpl, err := template.ParseGlob(HTML_PATH)
	if err != nil {
		log.Err(err).C("Error setting up templates")
		return nil
	}

	return &service{
		logger:     log,
		markdowner: md,
		templates:  tpl,
	}
}
