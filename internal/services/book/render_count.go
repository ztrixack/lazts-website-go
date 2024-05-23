package book

import (
	"fmt"
	"io"
)

func (s *service) RenderCount(wr io.Writer) error {
	count, err := getCount("books")
	if err != nil {
		s.logger.Err(err).E("Error getting book count")
		return err
	}

	s.logger.Fields("count", count).I("book count")

	if _, err := wr.Write([]byte(fmt.Sprintf("%d", count))); err != nil {
		s.logger.Err(err).E("Error executing book count template")
		return err
	}

	return nil
}
