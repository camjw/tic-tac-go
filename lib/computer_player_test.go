package lib

import "testing"

func TestAvoidsLosses(t *testing.T) {
	board := NewBoard("X")
	for i := 0; i < 2; i++ {
		board.PlayMove(2 * i, "X")
	}
	computer := NewComputer(board)
	computer.ComputerMove()
	got := computer.NextMove
	want := 1
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTakesWinners(t *testing.T) {
	board := NewBoard("X")
	for _, move := range [3]int{0, 6, 7} {
		board.PlayMove(move, "X")
	}
	for _, move := range [2]int{3, 4} {
		board.PlayMove(move, "O")
	}
	computer := NewComputer(board)
	computer.ComputerMove()
	got := computer.NextMove
	want := 5
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIndexInSliceWhenNotElement(t *testing.T) {
	got := indexInSlice([]float64{1.0}, 0.0)
	want := -1
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
