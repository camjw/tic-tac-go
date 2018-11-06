package lib

import "testing"

func TestPlayerMove(t *testing.T) {
	player := NewPlayer(&mockboard)
	player.PlayerMove(0)
	got := len(mockboard.PlayMoveCalls())
	want := 1
	if got != want {
		t.Errorf("Expected %v calls but got %v calls", got, want)
	}
}

func TestPlayerMoveSymbol(t *testing.T) {
	player := NewPlayer(&mockboard)
	player.PlayerMove(1)
	got := player.Symbol
	want := "X"
	if got != want {
		t.Errorf("got %.s want %.s", got, want)
	}
}
