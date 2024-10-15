package game

import (
	"context"
	"math"
)

type BestMove struct {
	Row, Col, score int
}

func newBestMove(row, col, score int) *BestMove {
	return &BestMove{Row: row, Col: col, score: score}
}

func minimax(ctx context.Context, board Board, player Symbol, depth int, isMax bool, alpha, beta int) int {
	score := board.evaluate(player)

	if score == 1 || score == -1 {
		return score
	}

	if board.IsFull() {
		return 0
	}

	var bestScore int

	if isMax {
		for _, cell := range board.emptyCells() {
			bestScore = math.MinInt
			i, j := cell.Row, cell.Col
			board[i][j] = player
			bestScore = _max(bestScore, minimax(ctx, board, player, depth+1, false, alpha, beta))
			board[i][j] = Empty
			alpha = _max(alpha, bestScore)

			if beta <= alpha {
				break
			}
		}

	} else {
		bestScore = math.MaxInt
		for _, cell := range board.emptyCells() {
			i, j := cell.Row, cell.Col
			board[i][j] = OpponentSymbol(player)
			bestScore = _min(bestScore, minimax(ctx, board, player, depth+1, true, alpha, beta))
			board[i][j] = Empty
			beta = _min(beta, bestScore)

			if beta <= alpha {
				break
			}
		}

	}

	if isTimeout(ctx) {
		return bestScore
	}

	return bestScore

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
