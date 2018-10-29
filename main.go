package main

import board "github.com/camjw/tic-tac-go/lib"
import "fmt"

func main() {
  var board = board.Board{}
  board.Initialize()
  board.PlayMove(0,"X")
  board.PlayMove(3,"X")
  board.PlayMove(6,"X")
  board.Print()
  fmt.Println(board.VerticalWin("X"))
}
