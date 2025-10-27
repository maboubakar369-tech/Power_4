package main

const (
	Rows    = 6
	Columns = 7
)

type Game struct {
	Board         [Rows][Columns]int
	CurrentPlayer int
	GameOver      bool
	Winner        int
}

func NewGame() *Game {
	return &Game{
		CurrentPlayer: 1,
	}
}

func (g *Game) Play(col int) bool {
	if g.GameOver || col < 0 || col >= Columns {
		return false
	}

	for row := Rows - 1; row >= 0; row-- {
		if g.Board[row][col] == 0 {
			g.Board[row][col] = g.CurrentPlayer
			if g.checkWin(row, col) {
				g.GameOver = true
				g.Winner = g.CurrentPlayer
			} else {
				g.switchPlayer()
			}
			return true
		}
	}
	return false
}

func (g *Game) switchPlayer() {
	if g.CurrentPlayer == 1 {
		g.CurrentPlayer = 2
	} else {
		g.CurrentPlayer = 1
	}
}

func (g *Game) checkWin(row, col int) bool {
	player := g.Board[row][col]
	dirs := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}
	for _, d := range dirs {
		count := 1
		count += g.count(row, col, d[0], d[1], player)
		count += g.count(row, col, -d[0], -d[1], player)
		if count >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) count(row, col, dRow, dCol, player int) int {
	count := 0
	for i := 1; i < 4; i++ {
		nr, nc := row+dRow*i, col+dCol*i
		if nr < 0 || nr >= Rows || nc < 0 || nc >= Columns {
			break
		}
		if g.Board[nr][nc] == player {
			count++
		} else {
			break
		}
	}
	return count
}

func (g *Game) Reset() {
	*g = *NewGame()
}
