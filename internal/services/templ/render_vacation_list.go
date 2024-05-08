package templ

import "io"

func (s *service) RenderVacationList(wr io.Writer) error {
	data, err := s.getVacationList()
	if err != nil {
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "vacation_list.html", data); err != nil {
		s.log.Err(err).E("Error executing vacations template")
		return err
	}
	return nil
}
