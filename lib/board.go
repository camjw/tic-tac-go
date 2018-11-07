package lib

import "fmt"

type Board struct {
	// Tic-Tac-Toe board struct.
	Grid       [3][3]string
	TotalMoves int
	NextTurn   string
}

func NewBoard(symbol string) *Board {
	// Returns a new instance of the Board Struct
	EmptyGrid := [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	return &Board{EmptyGrid, 0, symbol}
}

func (b Board) WhoseTurn() string {
	// Checks whose turn it currently is
	return b.NextTurn
}

func (b Board) Print() {
	// Prints the board to Stdout
	for i := range b.Grid {
		fmt.Println(b.Grid[i])
	}
}

//
func (b *Board) switchPlayer() {
	if b.WhoseTurn() == "X" {
		b.NextTurn = "O"
	} else {
		b.NextTurn = "X"
	}
}

func (b *Board) PlayMove(index int, symbol string) {
	// Plays a new move on the grid
	row, column := index/3, index%3
	b.Grid[row][column] = symbol
	b.TotalMoves += 1
	b.switchPlayer()
}

func (b *Board) GetValidMoves() []int {
	// Returns the currently valid moves
	output := []int{}
	for i := range b.Grid {
		for j := range b.Grid[i] {
			if b.Grid[i][j] == "-" {
				output = append(output, 3*i+j)
			}
		}
	}
	if b.GameOver() {
		return []int{}
	}
	return output
}

func (b Board) Winner(symbol string) bool {
	// Checks if the game has been won by the passed symbol
	if b.verticalWin(symbol) || b.horizontalWin(symbol) || b.diagonalWin(symbol) {
		return true
	}
	return false
}

func (b Board) GameOver() bool {
	// Checks if the game is over (either player won or a draw)
	if b.Winner("X") || b.Winner("O") || b.TotalMoves == 9 {
		return true
	}
	return false
}

func (b Board) verticalWin(symbol string) bool {
	for i := 0; i < 3; i++ {
		column := [3]string{b.Grid[0][i], b.Grid[1][i], b.Grid[2][i]}
		if column == [3]string{symbol, symbol, symbol} {
			return true
		}
	}
	return false
}

func (b Board) horizontalWin(symbol string) bool {
	for i := 0; i < 3; i++ {
		row := b.Grid[i]
		if row == [3]string{symbol, symbol, symbol} {
			return true
		}
	}
	return false
}

func (b Board) diagonalWin(symbol string) bool {
	ddiag := [3]string{b.Grid[0][0], b.Grid[1][1], b.Grid[2][2]}
	udiag := [3]string{b.Grid[0][2], b.Grid[1][1], b.Grid[2][0]}
	winner := [3]string{symbol, symbol, symbol}
	if udiag == winner || ddiag == winner {
		return true
	}
	return false
}

func (b Board) Clone() Board {
	// Returns a clone of the current board
	return b
}
