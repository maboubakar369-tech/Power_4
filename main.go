package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var game = NewGame()

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", serveHome)     // Page d'accueil
	http.HandleFunc("/game", serveGame) // Page du jeu
	http.HandleFunc("/move", handleMove)
	http.HandleFunc("/reset", handleReset)

	log.Println("✅ Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Page d'accueil
func serveHome(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/homes.html"))
	tpl.Execute(w, nil)
}

// Page du jeu
func serveGame(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct {
		Game *Game
		Cols []int
	}{
		Game: game,
		Cols: []int{0, 1, 2, 3, 4, 5, 6},
	}
	tpl.Execute(w, data)
}

// Jouer un coup
func handleMove(w http.ResponseWriter, r *http.Request) {
	colStr := r.URL.Query().Get("col")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Error(w, "invalid column", http.StatusBadRequest)
		return
	}

	if !game.GameOver {
		game.Play(col)
	}

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// Réinitialiser la partie
func handleReset(w http.ResponseWriter, r *http.Request) {
	game = NewGame()
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
