package notes

import (
	"fmt"
	"net/http"
	"strings"
)

func (h *handler) Content(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := h.routeContent(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (h *handler) routeContent(r *http.Request, w http.ResponseWriter) error {
	path := strings.TrimPrefix(r.URL.Path, "/_notes/contents/")
	parts := strings.Split(path, "/")

	switch len(parts) {
	case 1:
		return h.pager.Render(w, "notes")
	case 2:
		return h.pager.RenderMarkdown(w, "notes", parts[1])

	default:
		return fmt.Errorf("page not found")
	}
}
