package page

import (
	"io"
)

type CloudData struct {
	Items []Cloud
}

func (s *service) RenderCloud(wr io.Writer, count int) error {
	if err := s.templates.ExecuteTemplate(wr, "cloud.html", CloudData{randomizeClouds(count)}); err != nil {
		s.log.Err(err).E("Error executing hero cloud template")
		return err
	}
	return nil
}
