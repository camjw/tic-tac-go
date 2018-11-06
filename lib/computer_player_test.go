package lib

import "testing"
import "fmt"

// import "reflect"

func TestScoreGameNotEnded(t *testing.T) {
	computer := NewComputer(&mockboard)
	got := computer.Score(&mockboard)
	want := 0.0
	if got != want {
		t.Errorf("Got %.f calls but want %.f", got, want)
	}
}

func TestScoreXWinner(t *testing.T) {
	board := NewBoard("X")
	for i := 0; i < 3; i++ {
		board.PlayMove(i, "X")
	}
	computer := NewComputer(board)
	got := computer.Score(board)
	want := -1.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestScoreOWinner(t *testing.T) {
	board := NewBoard("X")
	for i := 0; i < 3; i++ {
		board.PlayMove(i, "O")
	}
	computer := NewComputer(board)
	got := computer.Score(board)
	want := 1.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

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
	fmt.Println(board.WhoseTurn())
	board.Print()
	computer := NewComputer(board)
	computer.ComputerMove()
	got := computer.NextMove
	want := 5
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
