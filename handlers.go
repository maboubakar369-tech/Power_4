package main

import (
	"html/template"
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	// Parse les templates base.html + home.html
	tmpl, err := template.ParseFiles("templates/base.html", "templates/home.html")
	if err != nil {
		log.Println("Erreur parsing template:", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}

	// Exécute le template "base"
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println("Erreur exécution template:", err)
	}
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	// Ici tu traiteras les mouvements envoyés par le frontend
	w.Write([]byte("Move reçu"))
}
