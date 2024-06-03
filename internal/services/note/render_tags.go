package note

import (
	"io"
)

type TagsData struct {
	Items []TagHTML
}

func (s *service) RenderTags(wr io.Writer) error {
	items, err := s.getTagList("notes")
	if err != nil {
		s.logger.Err(err).E("Error getting tag list")
		return err
	}

	s.logger.Fields("count", len(items)).I("rendered tag notes")

	if err := s.templates.ExecuteTemplate(wr, "tags.html", TagsData{items}); err != nil {
		s.logger.Err(err).E("Error executing note tags template")
		return err
	}
	return nil
}
