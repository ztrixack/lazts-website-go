package system

import "net/http"

func (h *handler) CSS() http.Handler {
	fs := http.FileServer(http.Dir("web/static/css"))
	return http.StripPrefix("/static/css", fs)
}
