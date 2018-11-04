package lib

type BoardToPlay interface {
  PlayMove(int, string)
  GetValidMoves() []int
}
