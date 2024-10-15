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

	opponent := OpponentSymbol(player)

	if isMax {
		best := math.MinInt

		for _, cell := range board.emptyCells() {
			i, j := cell.Row, cell.Col

			board[i][j] = player
			best = _max(best, minimax(ctx, board, player, depth+1, false, alpha, beta))
			board[i][j] = Empty

			alpha = _max(alpha, best)

			if beta <= alpha {
				break
			}

			if isTimeout(ctx) {
				break
			}
		}

		return best

	} else {

		best := math.MaxInt

		for _, cell := range board.emptyCells() {
			i, j := cell.Row, cell.Col

			board[i][j] = opponent
			best = _min(best, minimax(ctx, board, player, depth+1, true, alpha, beta))
			board[i][j] = Empty

			beta = _min(beta, best)

			if beta <= alpha {
				break
			}

			if isTimeout(ctx) {
				break
			}
		}

		return best

	}
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
