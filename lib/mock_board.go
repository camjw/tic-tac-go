package lib

type MockBoard struct {
  Index int
  Symbol string
}

func (m *MockBoard) PlayMove(index int, symbol string) () {
  m.Index = index
  m.Symbol = symbol
}

func (m MockBoard) GetValidMoves() ([]int) {
  return []int{0}
}

func (m MockBoard) Winner(symbol string) (bool) {
  return false
}

func (m MockBoard) GameOver() (bool) {
  return false
}
