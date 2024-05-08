package vacations

import (
	"net/http"
	"strings"
)

func (h *handler) Content(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := h.routeContent(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (h *handler) routeContent(r *http.Request, w http.ResponseWriter) error {
	path := strings.TrimPrefix(r.URL.Path, "/_vacations/contents/")
	parts := strings.Split(path, "/")

	return h.hs.RenderMarkdown(w, "vacations", parts[0])
}
