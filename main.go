package main

import lib "github.com/camjw/tic-tac-go/lib"
import "fmt"

func main() {
  var board = lib.NewBoard()
  var player = lib.NewPlayer("X", board)
  player.PlayMove(0)
  player.PlayMove(3)
  player.PlayMove(6)
  board.Print()
  fmt.Println(board.VerticalWin("X"))
  fmt.Println(board.TotalMoves)
}
