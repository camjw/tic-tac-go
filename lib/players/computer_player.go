package players

import "sort"
import "fmt"

type ComputerPlayer struct {
  Symbol string
  Board BoardToPlay
  NextMove int
}

func NewComputer(board BoardToPlay) (*ComputerPlayer) {
  return &ComputerPlayer{"O", board, -1}
}

func (c *ComputerPlayer) ComputerMove() () {
  c.MiniMax(c.Board)
  c.Board.PlayMove(c.NextMove, c.Symbol)
}


func (c ComputerPlayer) GetValidMoves() ([]int) {
  return c.Board.GetValidMoves()
}

func (c ComputerPlayer) Score(game BoardToPlay) (float64) {
  if (game.Winner("O")) {
    return 1.0
  } else if (game.Winner("X")) {
    return -1.0
  } else {
    return 0.0
  }
}

func (c ComputerPlayer) MiniMax(board BoardToPlay) (float64) {

}

func minFloat64Slice(slice []float64) (float64) {
  sort.Float64s(slice)
  return slice[0]
}

func maxFloat64Slice(slice []float64) (float64) {
  sort.Float64s(slice)
  return slice[len(slice) - 1]
}

func indexInSlice(slice []float64, value float64) (int) {
  for i, v := range slice {
      if (v == value) {
          return i
      }
  }
  return -1
}
