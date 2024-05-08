package templ

import (
	"io"
)

type VacationHighlightData struct {
	Title    string
	Excerpt  string
	Image    string
	Link     string
	ShowDate string
	Location string
}

func (s *service) RenderVacationHighlight(wr io.Writer) error {
	data, err := s.getVacation()
	if err != nil {
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "vacation_highlight.html", data); err != nil {
		s.log.Err(err).E("Error executing vacation highlight template")
		return err
	}
	return nil
}
