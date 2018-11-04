package players

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

func (c *ComputerPlayer) MiniMax(board BoardToPlay) (float64) {
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

  if (c.Board.WhoseTurn == "O") {
    max_score_
  }
  // if game.active_turn == @player
  //       # This is the max calculation
  //       max_score_index = scores.each_with_index.max[1]
  //       @choice = moves[max_score_index]
  //       return scores[max_score_index]
  //   else
  //       # This is the min calculation
  //       min_score_index = scores.each_with_index.min[1]
  //       @choice = moves[min_score_index]
  //       return scores[min_score_index]
  //   end

  return 1.5
}

func minFloat64Slice(s []float64) int {
  sort.Float64s(s)
  return s[0]
}

func maxFloat64Slice(s []float64) int {
  sort.Float64s(s)
  return s[len(s) - 1]
}
