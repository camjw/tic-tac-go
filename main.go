package main

import board "github.com/camjw/tic-tac-go/lib/board"
import players "github.com/camjw/tic-tac-go/lib/players"

import "fmt"

func main() {
  board := board.New("X")
  player := players.NewPlayer(board)
  computer := players.NewComputer(board)
  player.PlayerMove(0)
  board.Print()
  fmt.Println(board.VerticalWin("X"))
  fmt.Println(board.TotalMoves)
  var possible_board struct
  for _, move := range board.GetValidMoves() {
    possible_board = board
    possible_board.PlayMove(move, "O")
    possible_board.Print()
  }
  // fmt.Println(computer.MiniMax(board))

  fmt.Println(computer.NextMove)
}
