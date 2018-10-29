package lib

import "fmt"

type Board struct {
  BoardState [][]string
}

func (b *Board) Initialize() () {
  b.BoardState = [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }
}

func (b Board) Print() () {
  for i, _ := range b.BoardState {
    fmt.Println(b.BoardState[i])
  }
}

func (b *Board) PlayMove(index_x int, index_y int, token string) () {
  b.BoardState[index_y][index_x] = token
}

func (b Board) GameOver() (bool, string) {
   
}

func (b Board) VerticalWin() (bool, string) {

}

func (b Board) HorizontalWin() (bool, string) {

}

func (b Board) DiagonalWin() (bool, string) {

}
