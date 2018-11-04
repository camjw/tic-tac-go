package lib

import "testing"

type MockBoard struct {
  Index int
  Symbol string
}

func (m *MockBoard) PlayMove(index int, symbol string) () {
  m.Index = index
  m.Symbol = symbol
}

func TestPlayerMove(t *testing.T) {
  mockboard := MockBoard{}
  player := NewPlayer("X", &mockboard)
  player.PlayerMove(1)
  got_index := mockboard.Index
  got_symbol := mockboard.Symbol
  want_index := 1
  want_symbol := "X"
  if (got_index != want_index) || (got_symbol != want_symbol) {
    t.Errorf("Either index or symbol is wrong")
  }
}
