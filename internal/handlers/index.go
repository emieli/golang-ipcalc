package handlers

import (
	"fmt"
	"ipcalc/internal/templates"
	"log"
	"net/http"
	"net/netip"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := templates.Hello("Emil").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func Prefix(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	input := r.PostForm.Get("prefix")
	input = strings.TrimSpace(input)
	prefix, err := netip.ParsePrefix(input)
	if err != nil {
		log.Println(err)
		msg := fmt.Sprintf("invalid prefix %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	err = templates.Prefix(prefix).Render(r.Context(), w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
