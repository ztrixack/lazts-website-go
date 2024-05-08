package home

import (
	"net/http"
	"strconv"
)

func (h *handler) Blackhole(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	count := 1000
	countStr := r.URL.Query().Get("count")
	if countStr != "" {
		cnt, err := strconv.Atoi(countStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		count = cnt
	}

	err := h.hs.RenderHeroBlackhole(w, count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
