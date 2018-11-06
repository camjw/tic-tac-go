package lib

import "sort"

type ComputerPlayer struct {
	Symbol   string
	Board    BoardToPlay
	NextMove int
}

func NewComputer(board BoardToPlay) *ComputerPlayer {
	return &ComputerPlayer{"O", board, -1}
}

func (c *ComputerPlayer) ComputerMove() {
	c.MiniMax(c.Board)
	c.Board.PlayMove(c.NextMove, c.Symbol)
}

func (c ComputerPlayer) Score(game BoardToPlay) float64 {
	if game.Winner("O") {
		return 1.0
	} else if game.Winner("X") {
		return -1.0
	} else {
		return 0.0
	}
}

func (c *ComputerPlayer) MiniMax(board BoardToPlay) float64 {
	if board.GameOver() {
		return c.Score(board)
	}

	scores := []float64{}
	moves := []int{}

	for _, move := range board.GetValidMoves() {
		possible_board := board.Clone()
		possible_board.PlayMove(move, possible_board.WhoseTurn())
		scores = append(scores, c.MiniMax(&possible_board))
		moves = append(moves, move)
	}

	return c.scorePosition(board, moves, scores)
}

func (c *ComputerPlayer) scorePosition(board BoardToPlay, moves []int, scores []float64) float64 {
	if board.WhoseTurn() == "O" {
		max_score_index := indexInSlice(scores, maxFloat64Slice(scores))
		c.NextMove = moves[max_score_index]
		return scores[max_score_index]
	} else {
		min_score_index := indexInSlice(scores, minFloat64Slice(scores))
		c.NextMove = moves[min_score_index]
		return scores[min_score_index]
	}
}

func minFloat64Slice(slice []float64) float64 {
	sort.Float64s(slice)
	return slice[0]
}

func maxFloat64Slice(slice []float64) float64 {
	sort.Float64s(slice)
	return slice[len(slice)-1]
}

func indexInSlice(slice []float64, value float64) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
