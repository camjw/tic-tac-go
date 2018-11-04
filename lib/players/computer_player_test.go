package players

import "testing"
import "reflect"

func TestComputerMoveIndex(t *testing.T) {
  mockboard := MockBoard{}
  computer := NewComputer(&mockboard)
  computer.ComputerMove()
  got := mockboard.Index
  want := 1
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestComputerMoveSymbol(t *testing.T) {
  mockboard := MockBoard{}
  computer := NewComputer(&mockboard)
  computer.ComputerMove()
  got := mockboard.Symbol
  want := "O"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestComputerGetValidMoves(t *testing.T) {
  mockboard := MockBoard{}
  computer := NewComputer(&mockboard)
  got := computer.GetValidBoardMoves()
  want := []int{0}
  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestScoreNotEnded( t *testing.T) {
  mockboard := MockBoard{}
  computer := NewComputer(&mockboard)
  got := computer.Score(&mockboard)
  want:= 0.0
  if got != want {
   t.Errorf("got %f want %f", got, want)
  }
}
