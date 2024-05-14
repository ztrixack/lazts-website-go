package system

import "net/http"

func (h *handler) Javascript() http.Handler {
	fs := http.FileServer(http.Dir("web/static/js"))
	return http.StripPrefix("/static/js", fs)
}
