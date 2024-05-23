package vacation

import (
	"io"
)

func (s *service) RenderHighlight(wr io.Writer) error {
	if err := s.templates.ExecuteTemplate(wr, "highlight.html", nil); err != nil {
		s.logger.Err(err).E("Error executing vacation highlight template")
		return err
	}
	return nil
}
