package handlers

import (
	"ipcalc/internal/templates"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	prefix, _ := strings.CutPrefix(r.URL.String(), "/")
	err := templates.Index(prefix).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	return
}
