package system

import (
	"net/http"
)

func (h *handler) Static(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fs := http.FileServer(http.Dir("static"))
	http.StripPrefix("/static/", fs).ServeHTTP(w, r)
}
