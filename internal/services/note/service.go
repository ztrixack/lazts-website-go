package note

import (
	"html/template"
	"io"
	"lazts/internal/modules/log"
	"lazts/internal/modules/md"
)

type Servicer interface {
	RenderHero(wr io.Writer) error
	RenderList(wr io.Writer) error
	RenderTags(wr io.Writer) error
}

type service struct {
	log       log.Moduler
	markdown  md.Moduler
	templates *template.Template
}

var _ Servicer = (*service)(nil)

const HTML_PATH = "templates/sections/notes/*.html"
const DOMAIN = "notes"

func New(log log.Moduler, md md.Moduler) *service {
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
