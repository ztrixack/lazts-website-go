package note

import (
	"io"
)

type ListData struct {
	Items []NoteHTML
}

func (s *service) RenderList(wr io.Writer) error {
	items, err := s.getList("notes")
	if err != nil {
		s.logger.Err(err).E("Error getting notes")
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "list.html", ListData{items}); err != nil {
		s.logger.Err(err).E("Error executing note list template")
		return err
	}
	return nil
}
