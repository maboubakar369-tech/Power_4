package main

import (
	"fmt"
	"net/http"

}
func main() {
    // sert /style/* (ex: /style/css.css)
    http.Handle("/style/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

    http.HandleFunc("/", serveHome)
    http.HandleFunc("/move", handleMove)

    log.Println("✅ Serveur lancé sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

hv v