package lib

import "testing"
import "reflect"

func TestNew(t *testing.T) {
	board := NewBoard("X")
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
	board := NewBoard("X")
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

func TestSwitchMoveX(t *testing.T) {
	board := NewBoard("X")
	board.SwitchPlayer()
	got := board.WhoseTurn()
	want := "O"
	if got != want {
		t.Errorf("got %.s want %.s", got, want)
	}
}

func TestSwitchMoveO(t *testing.T) {
	board := NewBoard("O")
	board.SwitchPlayer()
	got := board.WhoseTurn()
	want := "X"
	if got != want {
		t.Errorf("got %.s want %.s", got, want)
	}
}

func TestPlayMoveIncreaseTotalMoves(t *testing.T) {
	board := NewBoard("X")
	board.PlayMove(3, "X")
	got := board.TotalMoves
	want := 1
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetValidMoves(t *testing.T) {
	board := NewBoard("X")
	got := board.GetValidMoves()
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleBoard_Print() {
	board := NewBoard("X")
	board.Print()
	// Output:
	// [- - -]
	// [- - -]
	// [- - -]
}

func TestHorizontalWinTrue(t *testing.T) {
	board := NewBoard("X")
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
	board := NewBoard("X")
	for i := 0; i < 3; i++ {
		board.PlayMove(3*i, "X")
	}
	got := board.VerticalWin("X")
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestDiagonalWinTrue(t *testing.T) {
	board := NewBoard("X")
	for i := 0; i < 3; i++ {
		board.PlayMove(4*i, "X")
	}
	got := board.DiagonalWin("X")
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestWinner(t *testing.T) {
	board := NewBoard("X")
	for i := 0; i < 3; i++ {
		board.PlayMove(4*i, "X")
	}
	got := board.Winner("X")
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestNoWin(t *testing.T) {
	board := NewBoard("X")
	got := board.Winner("X")
	want := false
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestGameOverTrue(t *testing.T) {
	board := NewBoard("X")
	for i := 0; i < 3; i++ {
		board.PlayMove(4*i, "X")
	}
	got := board.GameOver()
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestGameOverFalse(t *testing.T) {
	board := NewBoard("X")
	got := board.GameOver()
	want := false
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestGameOverTooManyMoves(t *testing.T) {
	board := NewBoard("X")
	for _, i := range [5]int{0, 1, 5, 6, 7} {
		board.PlayMove(i, "X")
	}
	for _, i := range [4]int{2, 3, 4, 8} {
		board.PlayMove(i, "O")
	}
	got := board.GameOver()
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
