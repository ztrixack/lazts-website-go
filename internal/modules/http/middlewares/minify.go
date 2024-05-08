package middlewares

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

func Minify() *minify.M {
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("text/css", css.Minify)

	return m
}
