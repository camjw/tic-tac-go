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

// func (c ComputerPlayer) MiniMax(board BoardToPlay) (float64) {
//   if board.GameOver() {
//     fmt.Println("GAME OVER")
//     board.Print()
//     return c.Score(board)
//   }
//
//   scores := []float64{}
//   moves := []int{}
//
//   fmt.Println("GET VALID MOVES")
//   fmt.Println(board.GetValidMoves())
//   for index, move := range board.GetValidMoves() {
//     PossibleBoard := board
//     PossibleBoard.PlayMove(move, board.WhoseTurn())
//     scores = append(scores, c.MiniMax(PossibleBoard))
//     moves = append(moves, move)
//     fmt.Println(index, move)
//     fmt.Println(moves)
//   }
//
//   if (c.Board.WhoseTurn() == "O") {
//     max_score := maxFloat64Slice(scores)
//     max_score_index := indexInSlice(scores, max_score)
//     c.NextMove = moves[max_score_index]
//     return scores[max_score_index]
//   } else {
//     min_score := minFloat64Slice(scores)
//     min_score_index := indexInSlice(scores, min_score)
//     c.NextMove = moves[min_score_index]
//     return scores[min_score_index]
//   }
// }

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
