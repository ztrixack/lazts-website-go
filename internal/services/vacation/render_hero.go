package vacation

import (
	"io"
)

type HeroData struct {
	Items []VacationHTML
}

func (s *service) RenderHero(wr io.Writer) error {
	items, err := s.getList("vacations")
	if err != nil {
		s.logger.Err(err).E("Error getting vacation list")
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "hero.html", HeroData{items}); err != nil {
		s.logger.Err(err).E("Error executing hero vacations template")
		return err
	}
	return nil
}
