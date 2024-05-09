package page

import (
	"html/template"
	"io"
	"lazts/internal/modules/md"
	"lazts/pkg/logger"
)

type Servicer interface {
	Render(wr io.Writer, page string) error
	RenderMarkdown(wr io.Writer, domain string, page string) error
	RenderBlackhole(wr io.Writer, counter int) error
	RenderCloud(wr io.Writer, counter int) error
}

type service struct {
	log       logger.Logger
	markdown  md.Moduler
	templates *template.Template
}

var _ Servicer = (*service)(nil)

const HTML_PATH = "templates/**/*.html"

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
