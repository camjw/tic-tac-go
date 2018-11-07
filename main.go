package main

import (
	lib "github.com/camjw/tic-tac-go/lib"
	"os"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("Welcome to Tic-Tac-Go!")

	board := lib.NewBoard(lib.GetFirstPlayer())
	player := lib.NewPlayer(board)
	computer := lib.NewComputer(board)

	for (!board.GameOver()) {
		if (board.WhoseTurn() == "X") {
			player.PlayerMove(lib.GetMove(board, os.Stdin))
		} else {
			fmt.Println("Computer's go!")
			computer.ComputerMove()
		}
		board.Print()
	}

	if board.Winner("X") {
	  fmt.Println("Well done, you win! This shouldn't be possible, please inform the developers.")
	} else if board.Winner("O") {
		fmt.Println("Unlucky, you lose.")
	} else {
		fmt.Println("It's a tie!")
	}
}
