package lib

type Player struct {
	Symbol string
	Board  BoardToPlay
}

func NewPlayer(board BoardToPlay) *Player {
	return &Player{"X", board}
}

func (p *Player) PlayerMove(index int) {
	p.Board.PlayMove(index, p.Symbol)
}
