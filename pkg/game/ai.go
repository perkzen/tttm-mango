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

	if isMax {
		bestMove := newBestMove(-1, -1, math.MinInt)
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] == Empty {
					board[i][j] = player
					move := minimax(ctx, board, player, depth+1, false, alpha, beta)
					board[i][j] = Empty

					if move.score > bestMove.score {
						bestMove = newBestMove(i, j, move.score)
					}

					alpha = max(alpha, bestMove.score)

					select {
					case <-ctx.Done():
						return bestMove
					default:
					}

					if beta <= alpha {
						break
					}
				}
			}
		}
		return bestMove
	} else {
		bestMove := newBestMove(-1, -1, math.MaxInt)
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] == Empty {
					board[i][j] = opponentSymbol(player)
					move := minimax(ctx, board, player, depth+1, true, alpha, beta)
					board[i][j] = Empty

					if move.score < bestMove.score {
						bestMove = newBestMove(i, j, move.score)
					}
					beta = min(beta, bestMove.score)

					select {
					case <-ctx.Done():
						return bestMove
					default:
					}

					if beta <= alpha {
						break
					}
				}
			}
		}
		return bestMove
	}
}
