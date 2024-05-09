package home

import (
	"net/http"
)

func (h *handler) Notes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	if err := h.noter.RenderHero(w);err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
