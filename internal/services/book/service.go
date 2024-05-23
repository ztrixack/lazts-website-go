package book

import (
	"html/template"
	"io"
	"lazts/internal/modules/log"
)

type Servicer interface {
	RenderHero(wr io.Writer) error
	RenderFilter(wr io.Writer, search, catalog, status string) error
	RenderList(wr io.Writer, search, catalog, status string) error
	RenderCount(wr io.Writer) error
}

type service struct {
	logger    log.Moduler
	templates *template.Template
}

var _ Servicer = (*service)(nil)

const HTML_PATH = "web/templates/sections/books/*.html"

func New(log log.Moduler) *service {
	tpl, err := template.ParseGlob(HTML_PATH)
	if err != nil {
		log.Err(err).C("Error setting up templates")
	}

	return &service{
		logger:    log,
		templates: tpl,
	}
}
