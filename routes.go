package main

import "net/http"

// registerRoutes ajoute les routes au mux passé en paramètre
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", serveHome)
	mux.HandleFunc("/move", handleMove)
}
