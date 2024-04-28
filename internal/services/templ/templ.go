package templ

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"lazts/pkg/logger"
)

type PageData struct {
	Title   string
	Content template.HTML
}

type Servicer interface {
	Render(wr io.Writer, page string, title string) error
}

type service struct {
	log       logger.Logger
	templates *template.Template
}

func New(log logger.Logger) *service {
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

func (s *service) Render(wr io.Writer, name string, title string) error {
	var buf bytes.Buffer

	page := fmt.Sprintf("%s.html", name)
	if err := s.templates.ExecuteTemplate(&buf, page, nil); err != nil {
		s.log.Err(err).E("Error executing template")
		return err
	}

	return s.templates.ExecuteTemplate(wr, "layout.html", PageData{
		Title:   title,
		Content: template.HTML(buf.String()),
	})
}
