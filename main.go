package main

import lib "github.com/camjw/tic-tac-go/lib"

import "fmt"

func main() {
	board := lib.NewBoard("O")
	player := lib.NewPlayer(board)
	computer := lib.NewComputer(board)
  // computer.MiniMax(board)
  computer.ComputerMove()
	// player.PlayerMove(0)
  player.PlayerMove(4)
  fmt.Println(computer.Board)

  // computer.MiniMax(board)
  computer.ComputerMove()
  board.Print()
  // computer.ComputerMove()
  // board.Print()
  // fmt.Println(computer.Board)
  // computer.ComputerMove()
  // board.Print()
	// // for _, move := range board.GetValidMoves() {
	//   possible_board := *board
	//   possible_board.PlayMove(move, "O")
	//   possible_board.Print()
	// }
	// fmt.Println(computer.MiniMax(board))

	fmt.Println(computer.NextMove)
}
