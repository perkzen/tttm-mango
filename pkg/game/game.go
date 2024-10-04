package game

import "math"

type Game struct {
	Gid   string
	Size  int
	Board Board
}

func NewGame(gid string, size int, moves string) *Game {
	board := movesToBoard(size, moves)
	return &Game{Gid: gid, Board: board, Size: size}
}

func (g *Game) GetBestMove(player Symbol) [2]int {

	bestScore := math.MinInt
	bestMove := [2]int{-1, -1}

	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			if g.Board[i][j] == Empty {
				g.Board[i][j] = player

				score := minmax(g.Board, 0, false)

				g.Board[i][j] = Empty

				if score > bestScore {
					bestScore = score
					bestMove = [2]int{i, j}
				}
			}
		}
	}

	return bestMove
}
