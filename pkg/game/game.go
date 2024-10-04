package game

import (
	"context"
	"math"
	"time"
)

type Game struct {
	Gid   string
	Size  int
	Board Board
}

func NewGame(gid string, size int, moves string) *Game {
	board := movesToBoard(size, moves)
	return &Game{Gid: gid, Board: board, Size: size}
}

func (g *Game) GetBestMove(player Symbol) *BestMove {
	bestMove := newBestMove(-1, -1, math.MinInt)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			if g.Board[i][j] == Empty {
				g.Board[i][j] = player
				move := minimax(ctx, g.Board, player, 0, false, math.MinInt, math.MaxInt)
				g.Board[i][j] = Empty

				if move.score > bestMove.score {
					bestMove = newBestMove(i, j, move.score)
				}
			}
		}
	}

	return bestMove
}
