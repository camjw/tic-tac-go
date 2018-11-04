package lib

type Player struct {
  Symbol string
  Board BoardToPlay
}

func NewPlayer(symbol string, board BoardToPlay) (*Player) {
  return &Player{symbol, board}
}

func (p *Player) PlayerMove(index int) () {
  p.Board.PlayMove(index, p.Symbol)
}
