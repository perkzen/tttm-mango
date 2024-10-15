package game

import "sync"

type TranspositionTable struct {
	m map[string]int
	sync.RWMutex
}

func newTranspositionTable() *TranspositionTable {
	return &TranspositionTable{m: make(map[string]int)}
}

func (t *TranspositionTable) cacheTransposition(board Board, score int) {
	for _, transposedBoard := range board.getTranspositions() {
		boardKey := transposedBoard.String()
		t.Lock()
		t.m[boardKey] = score
		t.Unlock()
	}
}

func (t *TranspositionTable) getCachedTransposition(board Board) (int, bool) {
	for _, transposedBoard := range board.getTranspositions() {
		boardKey := transposedBoard.String()
		t.RLock()
		if cachedScore, found := t.m[boardKey]; found {
			t.RUnlock()
			return cachedScore, true
		}
		t.RUnlock()
	}
	return 0, false
}

var transpositionTable = newTranspositionTable()
