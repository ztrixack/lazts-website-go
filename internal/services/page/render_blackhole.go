package page

import (
	"io"
	"time"
)

type BlackholeData struct {
	Title    string
	Subtitle string
	Since    string
	Year     int
	Items    []Blackhole
}

func (s *service) RenderBlackhole(wr io.Writer, count int) error {
	data := BlackholeData{
		Title:    "lazts",
		Subtitle: "Event horizon of my knowledge",
		Since:    "Since 1991",
		Year:     time.Now().Year(),
		Items:    randomizeBlackholes(count),
	}

	if err := s.templates.ExecuteTemplate(wr, "blackhole.html", data); err != nil {
		s.log.Err(err).E("Error executing hero blackhole template")
		return err
	}
	return nil
}
