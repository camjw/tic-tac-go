package lib

type ComputerPlayer struct {
  Symbol string
  Board BoardToPlay
}

func NewComputer(symbol string, board BoardToPlay) (*ComputerPlayer) {
  return &ComputerPlayer{symbol, board}
}

func (c *ComputerPlayer) ComputerMove() () {
  index := c.DecideMove()
  c.Board.PlayMove(index, c.Symbol)
}

func (c ComputerPlayer) GetOtherSymbol() (string) {
  if (c.Symbol == "X") {
    return "O"
  }
  return "X"
}

func (c ComputerPlayer) GetValidBoardMoves() ([]int) {
  return c.Board.GetValidMoves()
}

func (c ComputerPlayer) Score(game BoardToPlay) (float64) {
  if (game.Winner(c.Symbol)) {
    return 1.0
  } else if (game.Winner(c.GetOtherSymbol())) {
    return -1.0
  } else {
    return 0.0
  }
}

func (c ComputerPlayer) DecideMove() (int) {
  return 1
}
