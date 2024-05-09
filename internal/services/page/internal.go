package page

import (
	"math"
	"math/rand/v2"
)

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
