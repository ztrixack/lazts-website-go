package page

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"
	"time"
)

type PageData struct {
	Title   string
	Year    int
	Content template.HTML
	Menu    []MenuItem
}

type MenuItem struct {
	Label string
	Path  string
}

func (s *service) Render(wr io.Writer, name string) error {
	var buf bytes.Buffer

	data := PageData{
		Title: "lazts",
		Year:  time.Now().Year(),
		Menu:  DEFAULT_MENU,
	}

	if err := s.templates.ExecuteTemplate(&buf, fmt.Sprintf("%s.html", name), data); err != nil {
		s.log.Err(err).E("Error executing page template")
		return err
	}

	data.Content = template.HTML(buf.String())

	var finalBuf bytes.Buffer
	if err := s.templates.ExecuteTemplate(&finalBuf, "layout.html", data); err != nil {
		s.log.Err(err).E("Error executing layout template")
		return err
	}

	htmlWithInlineCSS := injectInlineCSS(finalBuf.String())

	if strings.Contains(name, "content") {
		htmlWithInlineCSS = injectMarkdownCSS(htmlWithInlineCSS)
	} else {
		htmlWithInlineCSS = removeMarkdownCSS(htmlWithInlineCSS)
	}

	_, err := wr.Write([]byte(htmlWithInlineCSS))
	if err != nil {
		s.log.Err(err).E("Error writing final content")
		return err
	}

	return nil
}
