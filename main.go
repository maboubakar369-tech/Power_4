package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type PageData struct {
	Game *Game
	Cols []int
}

var game = NewGame()

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/move", handleMove)
	http.HandleFunc("/reset", handleReset)

	log.Println("✅ Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{
		Game: game,
		Cols: []int{0, 1, 2, 3, 4, 5, 6},
	}
	tpl.Execute(w, data)
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	colStr := r.URL.Query().Get("col")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	game.Play(col)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	game.Reset()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
