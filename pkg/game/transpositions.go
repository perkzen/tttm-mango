package game

type TranspositionTable struct {
	cache map[string]int
}

func newTranspositionTable() *TranspositionTable {
	return &TranspositionTable{cache: make(map[string]int)}
}

func (t *TranspositionTable) Lookup(b Board) (int, bool) {
	for _, transposition := range b.generateTranspositions() {
		if score, ok := t.cache[transposition.String()]; ok {
			return score, true
		}
	}
	return 0, false
}

func (t *TranspositionTable) Store(b Board, score int) {
	for _, transposition := range b.generateTranspositions() {
		t.cache[transposition.String()] = score
	}
}

var transpositionTable = newTranspositionTable()
