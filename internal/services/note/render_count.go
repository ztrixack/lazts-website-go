package note

import (
	"fmt"
	"io"
)

func (s *service) RenderCount(wr io.Writer) error {
	count, err := getCount("notes")
	if err != nil {
		s.logger.Err(err).E("Error getting note count")
		return err
	}

	s.logger.Fields("count", count).I("note count")

	if _, err := wr.Write([]byte(fmt.Sprintf("%d", count))); err != nil {
		s.logger.Err(err).E("Error executing note count template")
		return err
	}

	return nil
}
