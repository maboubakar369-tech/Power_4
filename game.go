package main

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

// Implémente checkWin selon ta logique
func (p *Game) checkWin(row, col int) bool {
	// TODO : ta logique de détection de victoire ici
	return false
}
