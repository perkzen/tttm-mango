package game

import "fmt"

type Symbol string

const (
	X     Symbol = "X"
	O     Symbol = "O"
	Empty Symbol = ""
)

func PlayerSymbol(s string) (Symbol, error) {
	switch s {
	case "X":
		return X, nil
	case "O":
		return O, nil
	default:
		return Empty, fmt.Errorf("invalid player symbol: %s", s)
	}
}

func opponentSymbol(player Symbol) Symbol {
	if player == X {
		return O
	}
	return X
}

func (s Symbol) String() string {
	return string(s)
}
