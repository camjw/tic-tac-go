package main

import lib "github.com/camjw/tic-tac-go/lib"

import "fmt"

func main() {
  board := lib.NewBoard("X")
  player := lib.NewPlayer(board)
  computer := lib.NewComputer(board)
  player.PlayerMove(0)
  // for _, move := range board.GetValidMoves() {
  //   possible_board := *board
  //   possible_board.PlayMove(move, "O")
  //   possible_board.Print()
  // }
  fmt.Println(computer.MiniMax(board))

  fmt.Println(computer.NextMove)
}
