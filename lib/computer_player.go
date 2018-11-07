package lib

type ComputerPlayer struct {
	// Simple AI for tic-tac-toe. Uses the minimax algorithm to choose moves
	Symbol   string
	Board    BoardToPlay
	NextMove int
}

func NewComputer(board BoardToPlay) *ComputerPlayer {
	// Returns a new instance of the ComputerPlayer
	return &ComputerPlayer{"O", board, -1}
}

func (c *ComputerPlayer) ComputerMove() {
	// Plays a move for the computer on the computers board
	c.miniMax(c.Board)
	c.Board.PlayMove(c.NextMove, c.Symbol)
}

func (c ComputerPlayer) score(game BoardToPlay) float64 {
	if game.Winner("O") {
		return 1.0
	} else if game.Winner("X") {
		return -1.0
	} else {
		return 0.0
	}
}

func (c *ComputerPlayer) miniMax(board BoardToPlay) float64 {
	if board.GameOver() {
		return c.score(board)
	}

	scores := []float64{}
	moves := []int{}

	for _, move := range board.GetValidMoves() {
		possibleBoard := board.Clone()
		possibleBoard.PlayMove(move, possibleBoard.WhoseTurn())
		scores = append(scores, 0.9*c.miniMax(&possibleBoard))
		moves = append(moves, move)
	}

	return c.scorePosition(board, moves, scores)
}

func (c *ComputerPlayer) scorePosition(board BoardToPlay, moves []int, scores []float64) float64 {
	if board.WhoseTurn() == "O" {
		maxScoreIndex := indexInSlice(scores, maxFloat64Slice(scores))
		c.NextMove = moves[maxScoreIndex]
		return scores[maxScoreIndex]
	}
	minScoreIndex := indexInSlice(scores, minFloat64Slice(scores))
	c.NextMove = moves[minScoreIndex]
	return scores[minScoreIndex]
}
