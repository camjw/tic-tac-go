package main

import lib "github.com/camjw/tic-tac-go/lib"
import "fmt"

func main() {
  board := lib.NewBoard()
  player := lib.NewPlayer(board)
  computer := lib.NewComputer(board)
  player.PlayerMove(0)
  board.Print()
  fmt.Println(board.VerticalWin("X"))
  fmt.Println(board.TotalMoves)
  fmt.Println(board.GetValidMoves())
  fmt.Println(computer.MiniMax(board))
}
