package note

import (
	"io"
)

type HeroData struct {
	Items []NoteHTML
}

func (s *service) RenderHero(wr io.Writer) error {
	items, err := s.getList("notes")
	if err != nil {
		s.log.Err(err).E("Error getting notes")
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "hero.html", HeroData{items}); err != nil {
		s.log.Err(err).E("Error executing hero notes template")
		return err
	}
	return nil
}
