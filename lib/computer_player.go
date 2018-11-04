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

func (c ComputerPlayer) DecideMove() (int) {
  return 1
}
