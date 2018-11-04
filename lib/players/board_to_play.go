package players

type BoardToPlay interface {
  PlayMove(int, string)
  GetValidMoves() []int
  Winner(string) bool
  GameOver() bool
  WhoseTurn() string
  Print()
}
