package lib

// Struct for the player. Lets the player play a move and that is it.
type Player struct {
	Symbol string
	Board  BoardToPlay
}

// Returns a new player instance
func NewPlayer(board BoardToPlay) *Player {
	return &Player{"X", board}
}

// Plays a move for the player on the Board
func (p *Player) PlayerMove(index int) {
	p.Board.PlayMove(index, p.Symbol)
}
