package lib

//go:generate moq -out board_to_play_moq_test.go . BoardToPlay

type BoardToPlay interface {
	PlayMove(int, string)
	GetValidMoves() []int
	Winner(string) bool
	GameOver() bool
	WhoseTurn() string
	Print()
	Clone() Board
}
