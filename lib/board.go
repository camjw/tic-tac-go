package lib

import "fmt"

type Board struct {
  Grid [3][3]string
  TotalMoves int
}

func NewBoard() (*Board) {
  EmptyGrid := [3][3]string{
    [3]string{"-", "-", "-"},
    [3]string{"-", "-", "-"},
    [3]string{"-", "-", "-"},
  }
  return &Board{EmptyGrid, 0}
}

func (b Board) Print() () {
  for i := range b.Grid {
    fmt.Println(b.Grid[i])
  }
}

func (b *Board) PlayMove(index int, token string) () {
  row, column := index / 3, index % 3
  b.Grid[row][column] = token
  b.TotalMoves++
}

func (b *Board) GetValidMoves() ([]int) {
  output := []int{}
  for i := range b.Grid {
    for j := range b.Grid[i] {
      if b.Grid[i][j] == "-" {
        output = append(output, 3 * i + j)
      }
    }
  }
  return output
}

func (b Board) Winner(symbol string) (bool) {
  if b.VerticalWin(symbol) || b.HorizontalWin(symbol) || b.DiagonalWin(symbol) {
    return true
  }
  return false
}

func (b Board) GameOver() (bool) {
  if (b.Winner("X") || b.Winner("O")) {
    return true
  }
  return false
}

func (b Board) VerticalWin(symbol string) (bool) {
  for i := 0; i < 3; i++ {
    column := [3]string{b.Grid[0][i], b.Grid[1][i], b.Grid[2][i]}
    if column == [3]string{symbol, symbol, symbol} {
      return true
    }
  }
  return false
}

func (b Board) HorizontalWin(symbol string) (bool) {
  for i := 0; i < 3; i++ {
    row := b.Grid[i]
    if row == [3]string{symbol, symbol, symbol} {
      return true
    }
  }
  return false
}

func (b Board) DiagonalWin(symbol string) (bool) {
  ddiag := [3]string{b.Grid[0][0], b.Grid[1][1], b.Grid[2][2]}
  udiag := [3]string{b.Grid[0][2], b.Grid[1][1], b.Grid[0][2]}
  winner := [3]string{symbol, symbol, symbol}
  if udiag == winner || ddiag == winner {
    return true
  }
  return false
}
