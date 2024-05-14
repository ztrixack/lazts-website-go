package vacation

import (
	"io"
)

type HighlightData struct {
	Title    string
	Excerpt  string
	Image    string
	Link     string
	ShowDate string
	Location string
}

func (s *service) RenderHighlight(wr io.Writer) error {
	item, err := s.getOne("vacations")
	if err != nil {
		s.logger.Err(err).E("Error getting vacation highlight")
		return err
	}

	data := HighlightData{
		Title:    item.Title,
		Excerpt:  item.Excerpt,
		Image:    item.Image,
		Link:     item.Link,
		ShowDate: item.ShowDate,
		Location: item.Location,
	}

	if err := s.templates.ExecuteTemplate(wr, "highlight.html", data); err != nil {
		s.logger.Err(err).E("Error executing vacation highlight template")
		return err
	}
	return nil
}
