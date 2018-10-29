package main

import board "github.com/camjw/tic-tac-go/lib"

func main() {
  var board = board.Board{}
  board.Initialize()
  board.PlayMove(0,0,"X")
  board.Print()
}
