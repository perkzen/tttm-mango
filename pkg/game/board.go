package game

import (
	"strconv"
	"strings"
)

type Result string

const (
	PlayerX Result = "X"
	PlayerO Result = "O"
	Tie     Result = "tie"
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

		symbol, err := PlayerSymbol(parts[0])
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

func (b Board) checkWinner3x3Board() Result {
	size := 3

	// check rows
	for i := 0; i < size; i++ {
		if b[i][0] == b[i][1] && b[i][1] == b[i][2] && b[i][0] != Empty {
			return Result(b[i][0].String())
		}
	}

	// check columns
	for i := 0; i < size; i++ {
		if b[0][i] == b[1][i] && b[1][i] == b[2][i] && b[0][i] != Empty {
			return Result(b[0][i].String())
		}
	}

	// check diagonals
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[0][0] != Empty {
		return Result(b[0][0].String())
	}

	if b[0][2] == b[1][1] && b[1][1] == b[2][0] && b[0][2] != Empty {
		return Result(b[0][2].String())
	}

	return Tie
}

func (b Board) checkWinnerLargeBoard() Result {
	winCondition := 4 // 4 in a Row to win
	size := len(b)

	// check rows
	for i := 0; i < size; i++ {
		for j := 0; j <= size-winCondition; j++ {
			if b[i][j] != Empty && b[i][j] == b[i][j+1] && b[i][j+1] == b[i][j+2] && b[i][j+2] == b[i][j+3] {
				return Result(b[i][j].String())
			}
		}
	}

	// check columns
	for i := 0; i <= size-winCondition; i++ {
		for j := 0; j < size; j++ {
			if b[i][j] != Empty && b[i][j] == b[i+1][j] && b[i+1][j] == b[i+2][j] && b[i+2][j] == b[i+3][j] {
				return Result(b[i][j].String())
			}
		}
	}

	// check diagonals (top-left to bottom-right)
	for i := 0; i <= size-winCondition; i++ {
		for j := 0; j <= size-winCondition; j++ {
			if b[i][j] != Empty && b[i][j] == b[i+1][j+1] && b[i+1][j+1] == b[i+2][j+2] && b[i+2][j+2] == b[i+3][j+3] {
				return Result(b[i][j])
			}
		}
	}

	// check diagonals (top-right to bottom-left)
	for i := 0; i <= size-winCondition; i++ {
		for j := winCondition - 1; j < size; j++ {
			if b[i][j] != Empty && b[i][j] == b[i+1][j-1] && b[i+1][j-1] == b[i+2][j-2] && b[i+2][j-2] == b[i+3][j-3] {
				return Result(b[i][j].String())
			}
		}
	}

	return Tie
}

func (b Board) checkWinner() Result {
	size := len(b)

	if size == 3 {
		winner := b.checkWinner3x3Board()
		return winner
	}

	return b.checkWinnerLargeBoard()
}

func (b Board) evaluate(player Symbol) int {
	winner := b.checkWinner()

	playerSymbol, _ := PlayerSymbol(player.String())
	opponentSymbol := OpponentSymbol(playerSymbol)

	if winner.String() == playerSymbol.String() {
		return 1
	}

	if winner.String() == opponentSymbol.String() {
		return -1
	}

	return 0
}

type Cell struct {
	Row, Col int
}

func (b Board) emptyCells() []Cell {
	var cells []Cell

	for i := range b {
		for j := range b[i] {
			if b[i][j] == Empty {
				cells = append(cells, Cell{i, j})
			}
		}
	}

	return cells
}

func (b Board) IsFull() bool {
	cells := b.emptyCells()
	return len(cells) == 0
}
