package templ

import (
	"io"
)

type VacationData struct {
	Items []VacationHTML
}

func (s *service) RenderHeroVacations(wr io.Writer) error {
	data, err := s.getVacationList()
	if err != nil {
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "hero_vacations.html", data); err != nil {
		s.log.Err(err).E("Error executing hero vacations template")
		return err
	}
	return nil
}
