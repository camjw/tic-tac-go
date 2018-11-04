package lib

import "testing"

func TestNewBoard(t *testing.T) {
  board := NewBoard()
  got := board.Grid
  want := [3][3]string{
    [3]string{"-", "-", "-"},
    [3]string{"-", "-", "-"},
    [3]string{"-", "-", "-"},
  }
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestPlayMove(t *testing.T) {
  board := NewBoard()
  board.PlayMove(3, "X")
  got := board.Grid
  want := [3][3]string{
    [3]string{"-", "-", "-"},
    [3]string{"X", "-", "-"},
    [3]string{"-", "-", "-"},
  }
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}

func TestPlayMoveIncreaseTotalMoves(t *testing.T) {
  board := NewBoard()
  board.PlayMove(3, "X")
  got := board.TotalMoves
  want := 1
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func ExampleBoard_Print() {
  board := NewBoard()
  board.Print()
  // Output:
  // [- - -]
  // [- - -]
  // [- - -]
}


func TestHorizontalWinTrue(t *testing.T) {
  board := NewBoard()
  for i := 0; i < 3; i++ {
    board.PlayMove(i, "X")
  }
  got := board.HorizontalWin("X")
  want := true
  if got != want {
    t.Errorf("got %t want %t", got, want)
  }
}

func TestVerticalWinTrue(t *testing.T) {
  board := NewBoard()
  for i := 0; i < 3; i++ {
    board.PlayMove(3 * i, "X")
  }
  got := board.VerticalWin("X")
  want := true
  if got != want {
    t.Errorf("got %t want %t", got, want)
  }
}

func TestDiagonalWinTrue(t *testing.T) {
  board := NewBoard()
  for i := 0; i < 3; i++ {
    board.PlayMove(4 * i, "X")
  }
  got := board.DiagonalWin("X")
  want := true
  if got != want {
    t.Errorf("got %t want %t", got, want)
  }
}

func TestWinner(t *testing.T) {
  board := NewBoard()
  for i := 0; i < 3; i++ {
    board.PlayMove(4 * i, "X")
  }
  got := board.Winner("X")
  want := true
  if got != want {
    t.Errorf("got %t want %t", got, want)
  }
}

func TestNoWin(t *testing.T) {
  board := NewBoard()
  got := board.Winner("X")
  want := false
  if got != want {
    t.Errorf("got %t want %t", got, want)
  }
}
