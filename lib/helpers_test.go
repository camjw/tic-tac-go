package lib

import (
  "testing"
  "strings"
  "math/rand"
)

func TestIntInSliceTrue(t *testing.T) {
  got := intInSlice(6, []int{5,6})
  want := true
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestIntInSliceFalse(t *testing.T) {
  got := intInSlice(1, []int{5,6})
  want := false
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestGetMove(t *testing.T) {
  got := GetMove(&mockboard, strings.NewReader("0"))
  want := 0
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestGetMoveFailure(t *testing.T) {
  got := GetMove(&mockboard, strings.NewReader("f"))
  want := -1
  if got != want {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestGetFirstPlayer( t *testing.T) {
  rand.Seed(0)
  got := GetFirstPlayer()
  want := "X"
  if got != want {
    t.Errorf("got %.s want %.s", got, want)
  }
}
