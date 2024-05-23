package vacation

import (
	"fmt"
	"io"
)

func (s *service) RenderCount(wr io.Writer) error {
	count, err := getCount("vacations")
	if err != nil {
		s.logger.Err(err).E("Error getting vacation count")
		return err
	}

	s.logger.Fields("count", count).I("vacation count")

	if _, err := wr.Write([]byte(fmt.Sprintf("%d", count))); err != nil {
		s.logger.Err(err).E("Error executing vacation count template")
		return err
	}
	return nil
}
