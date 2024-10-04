package game

import "math"

var score = map[Result]int{
	WinX: 1,
	WinO: -1,
	Tie:  0,
}

func minmax(board Board, depth int, isMaximizing bool) int {
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
					score := minmax(board, depth+1, false)
					board[i][j] = Empty
					bestScore = max(bestScore, score)
				}
			}
		}
		return bestScore
	}

	bestScore := math.MaxInt
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == Empty {
				board[i][j] = O
				score := minmax(board, depth+1, true)
				board[i][j] = Empty
				bestScore = min(bestScore, score)
			}
		}
	}

	return bestScore
}
