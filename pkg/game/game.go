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

type BestMove struct {
	Row, Col, Score int
}

func newBestMove(row, col, score int) *BestMove {
	return &BestMove{Row: row, Col: col, Score: score}
}

func (g *Game) GetBestMove(player Symbol) *BestMove {
	bestMove := newBestMove(-1, -1, math.MinInt)

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	for _, cell := range g.Board.emptyCells() {
		i, j := cell.Row, cell.Col
		g.Board[i][j] = player
		score := minimax(ctx, g.Board, player, false, 0, math.MinInt, math.MaxInt)
		g.Board[i][j] = Empty

		if score > bestMove.Score {
			bestMove = newBestMove(i, j, score)
		}
	}

	return bestMove
}
