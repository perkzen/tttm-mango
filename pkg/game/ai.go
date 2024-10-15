package game

import (
	"context"
	"math"
)

func minimax(ctx context.Context, board Board, player Symbol, depth int, isMax bool, alpha, beta int) int {
	if cachedScore, found := transpositionTable.getCachedTransposition(board); found {
		return cachedScore
	}

	if score := board.evaluate(player); score == 1 || score == -1 {
		transpositionTable.cacheTransposition(board, score)
		return score
	}

	if board.IsFull() {
		transpositionTable.cacheTransposition(board, 0)
		return 0
	}

	if isMax {
		return maximizeMove(ctx, board, player, alpha, beta)
	} else {
		return minimizeMove(ctx, board, player, alpha, beta)
	}
}

func maximizeMove(ctx context.Context, board Board, player Symbol, alpha, beta int) int {
	bestScore := math.MinInt
	for _, cell := range board.emptyCells() {
		applyMove(board, cell, player)

		bestScore = max(bestScore, minimax(ctx, board, player, 0, false, alpha, beta))
		undoMove(board, cell)

		alpha = max(alpha, bestScore)
		if beta <= alpha || isTimeout(ctx) {
			break
		}
	}
	transpositionTable.cacheTransposition(board, bestScore)
	return bestScore
}

func minimizeMove(ctx context.Context, board Board, player Symbol, alpha, beta int) int {
	bestScore := math.MaxInt
	for _, cell := range board.emptyCells() {
		applyMove(board, cell, OpponentSymbol(player))

		bestScore = min(bestScore, minimax(ctx, board, player, 0, true, alpha, beta))
		undoMove(board, cell)

		beta = min(beta, bestScore)
		if beta <= alpha || isTimeout(ctx) {
			break
		}
	}
	transpositionTable.cacheTransposition(board, bestScore)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
