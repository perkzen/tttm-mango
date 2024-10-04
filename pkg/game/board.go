package game

import (
	"strconv"
	"strings"
)

type Result string

const (
	WinX Result = "X"
	WinO Result = "O"
	Tie  Result = "tie"
)

func (r Result) String() string {
	return string(r)
}

type Board [][]Symbol

func movesToBoard(size int, moves string) Board {
	board := newEmptyBoard(size)

	moveList := strings.Split(moves, "_")

	for _, move := range moveList {
		parts := strings.Split(move, "-")

		if len(parts) != 3 {
			continue
		}

		symbol, err := ToPlayerSymbol(parts[0])
		if err != nil {
			continue
		}

		row, rowErr := strconv.Atoi(parts[1])
		col, colErr := strconv.Atoi(parts[2])

		if rowErr != nil || colErr != nil || row < 0 || row >= size || col < 0 || col >= size {
			continue
		}

		board[row][col] = symbol
	}

	return board
}

func newEmptyBoard(size int) Board {
	board := make(Board, size)

	for i := range board {
		board[i] = make([]Symbol, size)
		for j := range board[i] {
			board[i][j] = Empty
		}
	}

	return board
}

func (board Board) checkWinner3x3Board() Result {
	size := 3

	// check rows
	for i := 0; i < size; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != Empty {
			return Result(board[i][0].String())
		}
	}

	// check columns
	for i := 0; i < size; i++ {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != Empty {
			return Result(board[0][i].String())
		}
	}

	// check diagonals
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != Empty {
		return Result(board[0][0].String())
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != Empty {
		return Result(board[0][2].String())
	}

	return Tie
}

func (board Board) checkWinnerLargeBoard() Result {
	winCondition := 4 // 4 in a Row to win
	size := len(board)

	// check rows
	for i := 0; i < size; i++ {
		for j := 0; j <= size-winCondition; j++ {
			if board[i][j] != Empty && board[i][j] == board[i][j+1] && board[i][j+1] == board[i][j+2] && board[i][j+2] == board[i][j+3] {
				return Result(board[i][j].String())
			}
		}
	}

	// check columns
	for i := 0; i <= size-winCondition; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] != Empty && board[i][j] == board[i+1][j] && board[i+1][j] == board[i+2][j] && board[i+2][j] == board[i+3][j] {
				return Result(board[i][j].String())
			}
		}
	}

	// check diagonals (top-left to bottom-right)
	for i := 0; i <= size-winCondition; i++ {
		for j := 0; j <= size-winCondition; j++ {
			if board[i][j] != Empty && board[i][j] == board[i+1][j+1] && board[i+1][j+1] == board[i+2][j+2] && board[i+2][j+2] == board[i+3][j+3] {
				return Result(board[i][j])
			}
		}
	}

	// check diagonals (top-right to bottom-left)
	for i := 0; i <= size-winCondition; i++ {
		for j := winCondition - 1; j < size; j++ {
			if board[i][j] != Empty && board[i][j] == board[i+1][j-1] && board[i+1][j-1] == board[i+2][j-2] && board[i+2][j-2] == board[i+3][j-3] {
				return Result(board[i][j].String())
			}
		}
	}

	return Tie
}

func (board Board) checkWinner() Result {
	size := len(board)

	if size == 3 {
		winner := board.checkWinner3x3Board()
		return winner
	}

	return board.checkWinnerLargeBoard()
}
