package notes

import (
	"net/http"
	"strings"
)

func (h *handler) Header(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/_notes/headers/")
	parts := strings.Split(path, "/")

	if err := h.noter.RenderHeader(w, parts[1]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
