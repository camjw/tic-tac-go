package lib

type ComputerPlayer struct {
  Symbol string
  Board BoardToPlay
  NextMove int
}

func NewComputer(board BoardToPlay) (*ComputerPlayer) {
  return &ComputerPlayer{"O", board, -1}
}

func (c *ComputerPlayer) ComputerMove() () {
  // index := c.MiniMax(c.Board)
  c.Board.PlayMove(1, c.Symbol)
}


func (c ComputerPlayer) GetValidBoardMoves() ([]int) {
  return c.Board.GetValidMoves()
}

func (c ComputerPlayer) Score(game BoardToPlay) (float64) {
  if (game.Winner("O")) {
    return 1.0
  } else if (game.Winner("X")) {
    return -1.0
  } else {
    return 0.0
  }
}

func (c ComputerPlayer) MiniMax(board BoardToPlay) (float64) {
  if c.Board.GameOver() {
    return c.Score(c.Board)
  }

  scores := []float64{}
  moves := []int{}

  for _, move := range c.GetValidBoardMoves() {
    PossibleBoard := c.Board
    PossibleBoard.PlayMove(move, "X")
    scores = append(scores, c.MiniMax(PossibleBoard))
    moves = append(moves, move)
  }

  return 1.5
}
