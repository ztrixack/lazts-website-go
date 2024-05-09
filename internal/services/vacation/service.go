package vacation

import (
	"html/template"
	"io"
	"lazts/internal/modules/md"
	"lazts/pkg/logger"
)

type Servicer interface {
	RenderHero(wr io.Writer) error
	RenderList(wr io.Writer) error
	RenderHighlight(wr io.Writer) error
}

type service struct {
	log       logger.Logger
	markdown  md.Moduler
	templates *template.Template
}

var _ Servicer = (*service)(nil)

const HTML_PATH = "templates/sections/vacations/*.html"

func New(log logger.Logger, md md.Moduler) *service {
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
