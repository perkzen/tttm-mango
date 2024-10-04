package game

import "math"

var score = map[Result]int{
	WinX: 1,
	WinO: -1,
	Tie:  0,
}

// minmax function with Alpha-Beta Pruning
func minmax(board Board, depth int, isMaximizing bool, alpha int, beta int) int {
	winner := checkWinner(board)

	if winner != Tie {
		return score[winner]
	}

	if isMaximizing {
		bestScore := math.MinInt
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] == Empty {
					board[i][j] = X
					score := minmax(board, depth+1, false, alpha, beta)
					board[i][j] = Empty
					bestScore = max(bestScore, score)
					alpha = max(alpha, bestScore)

					// Alpha-Beta Pruning
					if beta <= alpha {
						break // Cut off the remaining branches
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := math.MaxInt
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] == Empty {
					board[i][j] = O
					score := minmax(board, depth+1, true, alpha, beta)
					board[i][j] = Empty
					bestScore = min(bestScore, score)
					beta = min(beta, bestScore)

					// Alpha-Beta Pruning
					if beta <= alpha {
						break // Cut off the remaining branches
					}
				}
			}
		}
		return bestScore
	}
}
