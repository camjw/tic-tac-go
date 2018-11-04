package main

import lib "github.com/camjw/tic-tac-go/lib"
import "fmt"

func main() {
  var board = lib.NewBoard()
  var player = lib.NewPlayer(board)
  player.PlayerMove(0)
  player.PlayerMove(3)
  player.PlayerMove(6)
  board.Print()
  fmt.Println(board.VerticalWin("X"))
  fmt.Println(board.TotalMoves)
  fmt.Println(board.GetValidMoves())
}
