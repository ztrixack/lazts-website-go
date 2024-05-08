package templ

import (
	"io"
)

func (s *service) RenderNoteList(wr io.Writer) error {
	data, err := s.getNoteList()
	if err != nil {
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "note_list.html", data); err != nil {
		s.log.Err(err).E("Error executing note list template")
		return err
	}
	return nil
}
