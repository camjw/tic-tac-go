package main

import (
	lib "github.com/camjw/tic-tac-go/lib"
	"os"
)

func main() {
	board := lib.NewBoard("X")
	player := lib.NewPlayer(board)
	computer := lib.NewComputer(board)
	for (!board.GameOver()) {
		if (board.WhoseTurn() == "X") {
			player.PlayerMove(lib.GetMove(board, os.Stdin))
		} else {
			computer.ComputerMove()
		}
		board.Print()
	}
}
