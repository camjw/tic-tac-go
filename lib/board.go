package lib

import "fmt"

type Board struct {
  Grid [3][3]string
}

func (b *Board) Initialize() () {
  b.Grid = [3][3]string{
    [3]string{"-", "-", "-"},
    [3]string{"-", "-", "-"},
    [3]string{"-", "-", "-"},
  }
}

func (b Board) Print() () {
  for i, _ := range b.Grid {
    fmt.Println(b.Grid[i])
  }
}

func (b *Board) PlayMove(index int, token string) () {
  row, column := index / 3, index % 3
  b.Grid[row][column] = token
}

func (b Board) GameOver() (bool, string) {
  return true, "Hello"
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
