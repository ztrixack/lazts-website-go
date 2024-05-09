package vacation

import "io"

type ListData struct {
	Items []VacationHTML
}

func (s *service) RenderList(wr io.Writer) error {
	items, err := s.getList("vacations")
	if err != nil {
		s.log.Err(err).E("Error getting vacations")
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "list.html", ListData{items}); err != nil {
		s.log.Err(err).E("Error executing vacations template")
		return err
	}
	return nil
}
