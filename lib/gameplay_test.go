package lib

type MockBoard struct {
  Index int
  Symbol string
}

func (m *MockBoard) PlayMove(index int, symbol string) () {
  m.Index = index
  m.Symbol = symbol
}
