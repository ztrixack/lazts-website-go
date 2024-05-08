package notes

import (
	"fmt"
	"net/http"
	"strings"
)

func (h *handler) Page(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := h.routePage(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (h *handler) routePage(r *http.Request, w http.ResponseWriter) error {
	path := strings.TrimPrefix(r.URL.Path, "/notes/")
	parts := strings.Split(path, "/")

	switch len(parts) {
	case 1:
		if parts[0] == "" {
			return h.hs.Render(w, "notes")
		} else {
			return h.hs.Render(w, "notes_group")
		}
	case 2:
		return h.hs.Render(w, "notes_content")

	default:
		return fmt.Errorf("page not found")
	}
}
