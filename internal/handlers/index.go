package handlers

import (
	"ipcalc/internal/templates"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := templates.Hello("Emil").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
