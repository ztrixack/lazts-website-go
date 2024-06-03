package page

import (
	"fmt"
	"lazts/internal/utils"
	"math"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
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

func injectInlineCSS(html string) string {
	const stylesheet = "<link href=\"/static/css/app.css\" rel=\"stylesheet\" />"
	const tailwindcss = "/*!tailwindcss v3.4.3 | MIT License | https://tailwindcss.com*/"
	appcss, err := os.ReadFile(utils.GetStaticDir("css", "app.css"))
	if err != nil {
		return html
	}

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	minified, err := m.Bytes("text/css", appcss)
	if err != nil {
		return html
	}

	return strings.Replace(
		html,
		stylesheet,
		fmt.Sprintf("<style>%s</style>", strings.TrimPrefix(string(minified), tailwindcss)),
		1,
	)
}

func injectMarkdownCSS(html string) string {
	const stylesheet = "<link href=\"/static/css/markdown.css\" rel=\"stylesheet\" />"
	appcss, err := os.ReadFile(utils.GetStaticDir("css", "markdown.css"))
	if err != nil {
		return html
	}

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	minified, err := m.Bytes("text/css", appcss)
	if err != nil {
		return html
	}

	return strings.Replace(
		html,
		stylesheet,
		fmt.Sprintf("<style>%s</style>", string(minified)),
		1,
	)
}

func removeMarkdownCSS(html string) string {
	return strings.Replace(
		html,
		"<link href=\"/static/css/markdown.css\" rel=\"stylesheet\" />",
		"",
		1,
	)
}
