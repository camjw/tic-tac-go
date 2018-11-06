package lib

import "testing"

// import "reflect"

func TestComputerMove(t *testing.T) {
	computer := NewComputer(&mockboard)
	got := computer.Score(&mockboard)
	want := 0.0
	if got != want {
		t.Errorf("Got %.f calls but want %.f", got, want)
	}
}

// func TestScoreNotEnded(t *testing.T) {
// 	mockboard := BoardToPlayMock{}
// 	computer := NewComputer(&mockboard)
// 	got := computer.Score(&mockboard)
// 	want := 0.0
// 	if got != want {
// 		t.Errorf("got %f want %f", got, want)
// 	}
// }
