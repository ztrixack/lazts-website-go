package templ

import (
	"io"
	"math"
	"math/rand/v2"
	"time"
)

type Blackhole struct {
	Size    int
	Rotate  int
	Opacity int
	Width   int
}

type BlackholeData struct {
	Title    string
	Subtitle string
	Since    string
	Year     int
	Items    []Blackhole
}

func (s *service) RenderHeroBlackhole(wr io.Writer, count int) error {
	blackholes := randomizeBlackholes(count)
	data := BlackholeData{
		Title:    "lazts",
		Subtitle: "Event horizon of my knowledge",
		Since:    "Since 1991",
		Year:     time.Now().Year(),
		Items:    blackholes,
	}

	if err := s.templates.ExecuteTemplate(wr, "hero_blackhole.html", data); err != nil {
		s.log.Err(err).E("Error executing hero blackhole template")
		return err
	}
	return nil
}

func randomizeBlackholes(count int) []Blackhole {
	var blackholes []Blackhole
	for i := 0; i < count; i++ {
		size := rand.IntN(360) + 180
		rotate := rand.IntN(360)
		opacity := int(math.Max(110-float64(size*100/450), 5))
		width := (size - 90) / 6

		blackholes = append(blackholes, Blackhole{Size: size, Rotate: rotate, Opacity: opacity, Width: width})
	}
	return blackholes
}
