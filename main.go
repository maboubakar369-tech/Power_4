package main

import (
	"html/template"
	"net/http"
)
func main() {

	http.Handle("/style/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/move", handleMove)

	log.Println("✅ Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request){
	tpl := template.Must(template.ParseFiles("teste.html"))
	_= tpl.Execute(w, nil)
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Move reçu"))}

const (
	Rows    = 6
	Columns = 7
)

type Game struct {
	Board         [Rows][Columns]int
	currentplayer int
	GameOver      bool
	winner        int
}

func NewGame() *Game {
	return &Game{currentplayer: 1}
}

func (p *Game) play(col int) bool {
	if col < 0 || col >= Columns || p.GameOver {
		return false
	}
	for row := Rows - 1; row >= 0; row-- {
		cellstate := p.Board[row][col]
		if cellstate == 0 {
			p.Board[row][col] = p.currentplayer
			if p.checkWin(row, col) {
				p.GameOver = true
				p.winner = p.currentplayer
			} else {
				p.changeturn()
			}
			return true
		}
	}
	return false
}

func (p *Game) changeturn() {
	if p.currentplayer == 1 {
		p.currentplayer = 2
	} else {
		p.currentplayer = 1
	}

}
func (p *Game) checkWin(row, col int) bool {
	currentplayer := p.Board[row][col]
	directions := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}

	for _, d := range directions {
		count := 1
		count += p.count(row, col, d[0], d[1], currentplayer)
		count += p.count(row, col, -d[0], -d[1], currentplayer)
		if count >= 4 {
			return true
		}
	}
	return false
}
func (p *Game) count(row, col, dRow, dCol, player int) int {
	count := 0
	for i := 1; i < 4; i++ {
		nr := row + dRow*i
		nc := col + dCol*i

		if nr < 0 || nr >= Rows || nc < 0 || nc >= Columns {
			break
		}

		if p.Board[nr][nc] == player {
			count++
		} else {
			break
		}
	}
	return count
}
