package players

import "testing"

func TestPlayerMoveIndex(t *testing.T) {
  mockboard := MockBoard{}
  player := NewPlayer(&mockboard)
  player.PlayerMove(1)
  got := mockboard.Index
  want := 1
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestPlayerMoveSymbol(t *testing.T) {
  mockboard := MockBoard{}
  player := NewPlayer(&mockboard)
  player.PlayerMove(1)
  got := player.Symbol
  want := "X"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}
