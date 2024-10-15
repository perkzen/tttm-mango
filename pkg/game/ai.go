package game

import (
	"context"
	"math"
)

var scores = map[string]int{
	"X":   1,
	"O":   -1,
	"tie": 0,
}

type BestMove struct {
	Row, Col, score int
}

func newBestMove(row, col, score int) *BestMove {
	return &BestMove{Row: row, Col: col, score: score}
}

func minimax(ctx context.Context, board Board, player Symbol, depth int, isMax bool, alpha, beta int) *BestMove {
	winner := board.checkWinner()

	if winner != Tie {
		return newBestMove(-1, -1, scores[winner.String()])
	}

	var bestMove *BestMove
	if isMax {
		bestMove = newBestMove(-1, -1, math.MinInt)
	} else {
		bestMove = newBestMove(-1, -1, math.MaxInt)
	}

	for _, cell := range board.emptyCells() {
		row, col := cell.Row, cell.Col

		board[row][col] = player
		move := minimax(ctx, board, opponentSymbol(player), depth+1, !isMax, alpha, beta)
		board[row][col] = Empty

		if isMax {
			if move.score > bestMove.score {
				bestMove = newBestMove(row, col, move.score)
			}
			alpha = _max(alpha, bestMove.score)
		} else {
			if move.score < bestMove.score {
				bestMove = newBestMove(row, col, move.score)
			}
			beta = _min(beta, bestMove.score)
		}

		if beta <= alpha {
			break
		}

		if isTimeout(ctx) {
			return bestMove
		}
	}

	return bestMove
}

func _max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isTimeout(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
