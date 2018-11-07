package lib

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strconv"
)

func GetFirstPlayer() (string, error) {
	possibleMoves := [2]string{"X", "O"}
	return possibleMoves[rand.Intn(2)], nil
}

func GetMove(board BoardToPlay, reader io.Reader) int {
	for i := 0; i < 5; i++ {
		move, err := strconv.Atoi(getUserInput(board, reader))
		if (err == nil) && (intInSlice(move, board.GetValidMoves())) {
			return move
			break
		}
	}
	return -1
}

func getUserInput(board BoardToPlay, reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	fmt.Println("Valid moves: ", board.GetValidMoves())
	fmt.Print("Choose move!\n")
	scanner.Scan()
	return scanner.Text()
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func minFloat64Slice(slice []float64) float64 {
	return orderFloat64Slice(slice)[0]
}

func maxFloat64Slice(slice []float64) float64 {
	return orderFloat64Slice(slice)[len(slice)-1]
}

func indexInSlice(slice []float64, value float64) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func orderFloat64Slice(slice []float64) []float64 {
	output := []float64{}
	for _, value := range slice {
		output = append(output, value)
	}
	sort.Float64s(output)
	return output
}
