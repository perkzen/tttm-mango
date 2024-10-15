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

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	emptyCells := g.Board.emptyCells()

	for _, cell := range emptyCells {
		i, j := cell.Row, cell.Col
		g.Board[i][j] = player
		move := minimax(ctx, g.Board, player, 0, false, math.MinInt, math.MaxInt)
		g.Board[i][j] = Empty

		if move.score > bestMove.score {
			bestMove = newBestMove(i, j, move.score)
		}
	}

	return bestMove
}
