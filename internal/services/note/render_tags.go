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
		s.log.Err(err).E("Error getting tag list")
		return err
	}

	if err := s.templates.ExecuteTemplate(wr, "tags.html", TagsData{items}); err != nil {
		s.log.Err(err).E("Error executing note tags template")
		return err
	}
	return nil
}
