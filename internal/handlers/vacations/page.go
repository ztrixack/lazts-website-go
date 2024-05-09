package vacations

import (
	"net/http"
	"strings"
)

func (h *handler) Page(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := h.routePage(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (h *handler) routePage(r *http.Request, w http.ResponseWriter) error {
	path := strings.TrimPrefix(r.URL.Path, "/vacations/")
	parts := strings.Split(path, "/")

	if parts[0] == "" {
		return h.pager.Render(w, "vacations")
	} else {
		return h.pager.Render(w, "vacations_content")
	}

}
