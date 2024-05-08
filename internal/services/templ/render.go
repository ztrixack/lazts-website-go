package templ

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"time"
)

type PageData struct {
	Title   string
	Year    int
	Content template.HTML
	Menu    []MenuItem
}

type MenuItem struct {
	Path  string
	Icon  string
	Label string
}

func (s *service) Render(wr io.Writer, name string) error {
	var buf bytes.Buffer

	data := PageData{
		Title: "lazts",
		Year:  time.Now().Year(),
		Menu:  DEFAULT_MENU,
	}

	page := fmt.Sprintf("%s.html", name)
	if err := s.templates.ExecuteTemplate(&buf, page, data); err != nil {
		s.log.Err(err).E("Error executing page template")
		return err
	}

	data.Content = template.HTML(buf.String())
	if err := s.templates.ExecuteTemplate(wr, "layout.html", data); err != nil {
		s.log.Err(err).E("Error executing layout template")
		return err
	}

	return nil
}
