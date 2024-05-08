package templ

import (
	"html/template"
	"io"
	"lazts/internal/modules/md"
	"lazts/pkg/logger"
)

type Servicer interface {
	Render(wr io.Writer, page string) error
	RenderMarkdown(wr io.Writer, domain string, page string) error
	RenderHeroBlackhole(wr io.Writer, counter int) error
	RenderHeroBooks(wr io.Writer) error
	RenderHeroCloud(wr io.Writer, counter int) error
	RenderHeroVacations(wr io.Writer) error
	RenderHeroNotes(wr io.Writer) error
	RenderVacationList(wr io.Writer) error
	RenderVacationHighlight(wr io.Writer) error
	RenderBookList(wr io.Writer, search, catalog, status string) error
	RenderBookFilter(wr io.Writer, search, catalog, status string) error
	RenderNoteList(wr io.Writer) error
	RenderNoteTags(wr io.Writer) error
}

type service struct {
	log       logger.Logger
	markdown  md.Moduler
	templates *template.Template
}

var _ Servicer = (*service)(nil)

func New(log logger.Logger, md md.Moduler) *service {
	tpl, err := setupTemplates()
	if err != nil {
		log.Err(err).C("Error setting up templates")
	}

	names := []string{}
	for _, tmpl := range tpl.Templates() {
		names = append(names, tmpl.Name())
	}
	log.Fields("names", names).D("Templates loaded")

	return &service{
		log:       log,
		markdown:  md,
		templates: tpl,
	}
}

func setupTemplates() (*template.Template, error) {
	templates, err := template.ParseGlob("templates/**/*.html")
	if err != nil {
		return nil, err
	}

	return templates, nil
}
