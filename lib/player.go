package lib

type Player struct {
	// Struct for the player. Lets the player play a move and that is it.
	Symbol string
	Board  BoardToPlay
}

func NewPlayer(board BoardToPlay) *Player {
	// Returns a new player instance
	return &Player{"X", board}
}

func (p *Player) PlayerMove(index int) {
	// Plays a move for the player on the Board
	p.Board.PlayMove(index, p.Symbol)
}
