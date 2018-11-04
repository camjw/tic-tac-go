package lib

import "testing"

func TestComputerMoveIndex(t *testing.T) {
  mockboard := MockBoard{}
  computer := NewComputer("O", &mockboard)
  computer.ComputerMove()
  got := mockboard.Index
  want := 1
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestComputerMoveSymbol(t *testing.T) {
  mockboard := MockBoard{}
  computer := NewComputer("O", &mockboard)
  computer.ComputerMove()
  got := mockboard.Symbol
  want := "O"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}
