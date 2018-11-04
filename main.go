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
  fmt.Println(board.GetValidMoves())
  fmt.Println(computer.MiniMax(board))
  fmt.Println(computer.NextMove)
}
