package system

import "net/http"

func (h *handler) Error(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, "Resource not allowed", http.StatusMethodNotAllowed)
}
