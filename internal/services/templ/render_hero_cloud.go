package templ

import (
	"io"
	"math/rand/v2"
)

type Cloud struct {
	Top    int
	Left   int
	Rotate int
}

type CloudData struct {
	Items []Cloud
}

func (s *service) RenderHeroCloud(wr io.Writer, count int) error {
	data := CloudData{
		Items: randomizeClouds(count),
	}

	if err := s.templates.ExecuteTemplate(wr, "hero_cloud.html", data); err != nil {
		s.log.Err(err).E("Error executing hero cloud template")
		return err
	}
	return nil
}

func randomizeClouds(count int) []Cloud {
	var clouds []Cloud
	for i := 0; i < count; i++ {
		top := rand.IntN(150) - 50
		left := rand.IntN(100)
		rotate := rand.IntN(360)

		clouds = append(clouds, Cloud{Top: top, Left: left, Rotate: rotate})
	}
	return clouds
}
