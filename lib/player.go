package lib

type BoardToPlay interface {
  PlayMove(int, string)
}

type Player struct {
  Symbol string
  Board BoardToPlay
}

func (p *Player) PlayerMove(index int) () {
  p.Board.PlayMove(index, p.Symbol)
}

func NewPlayer(symbol string, board BoardToPlay) (*Player) {
	return &Player{symbol, board}
}
