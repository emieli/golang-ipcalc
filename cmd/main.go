package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	router.HandleFunc("GET /", getIndex)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
