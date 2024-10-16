package game

import (
	"context"
	"math"
)

func minimax(ctx context.Context, board Board, player Symbol, isMax bool, depth, alpha, beta int) int {
	if score := board.evaluate(player); score == 1 || score == -1 {
		return score
	}

	if board.IsFull() {
		return 0
	}

	if isMax {
		return maximizeMove(ctx, board, player, depth, alpha, beta)
	} else {
		return minimizeMove(ctx, board, player, depth, alpha, beta)
	}
}

func maximizeMove(ctx context.Context, board Board, player Symbol, depth, alpha, beta int) int {
	bestScore := math.MinInt
	for _, cell := range board.emptyCells() {
		applyMove(board, cell, player)

		bestScore = _max(bestScore, minimax(ctx, board, player, false, depth+1, alpha, beta))
		undoMove(board, cell)

		alpha = _max(alpha, bestScore)
		if beta <= alpha || isTimeout(ctx) {
			break
		}
	}
	return bestScore
}

func minimizeMove(ctx context.Context, board Board, player Symbol, depth, alpha, beta int) int {
	bestScore := math.MaxInt
	for _, cell := range board.emptyCells() {
		applyMove(board, cell, OpponentSymbol(player))

		bestScore = _min(bestScore, minimax(ctx, board, player, true, depth+1, alpha, beta))
		undoMove(board, cell)

		beta = _min(beta, bestScore)
		if beta <= alpha || isTimeout(ctx) {
			break
		}
	}
	return bestScore
}

func applyMove(board Board, cell Cell, player Symbol) {
	board[cell.Row][cell.Col] = player
}

func undoMove(board Board, cell Cell) {
	board[cell.Row][cell.Col] = Empty
}

func isTimeout(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
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
