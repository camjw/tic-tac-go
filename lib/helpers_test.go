package lib

import "testing"

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

}
