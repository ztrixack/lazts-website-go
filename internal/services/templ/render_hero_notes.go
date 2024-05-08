package templ

import (
	"io"
)

type NoteData struct {
	Items []NoteHTML
}

func (s *service) RenderHeroNotes(wr io.Writer) error {
	data, err := s.getNoteList()
	if err != nil {
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "hero_notes.html", data); err != nil {
		s.log.Err(err).E("Error executing hero notes template")
		return err
	}
	return nil
}
