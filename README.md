# Tic-Tac-Go
A tic-tac-toe game written in Go.
[![Build Status](https://travis-ci.com/camjw/tic-tac-go.svg?branch=master)](https://travis-ci.com/camjw/tic-tac-go)

![https://github.com/jpoles1/gopherbadger](./coverage_badge.png)

[![Maintainability](https://api.codeclimate.com/v1/badges/8ec28b43d94ce48ff7c9/maintainability)](https://codeclimate.com/github/camjw/tic-tac-go/maintainability)

[![Go Report Card](https://goreportcard.com/badge/github.com/camjw/tic-tac-go)](https://goreportcard.com/report/github.com/camjw/tic-tac-go)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Motivation

Tic-tac-toe has always been one of my favourite games, but one I rarely won at. I thought I would capture the tactical and strategic complexity of this game in a little Go app. The AI uses the MiniMax algorithm to choose its moves so, unfortunately, I am yet to beat this version too.

Please let me know if you win.

## Usage
### Installation

To install the app follow these steps from a directory including your `$GOPATH` (GOPATH explanation [here](https://golang.org/doc/code.html)):

```
$ git clone https://github.com/camjw/tic-tac-go.git
$ cd tic-tac-go
$ go install .
```

This will create an executable `tic-tac-go` which you can then run to start a game of Tic-Tac-Toe!

### Testing

To run the tests for this repo run the following command:
```
$ go test ./lib --cover
$ #=> ok  	github.com/camjw/tic-tac-go/lib	0.015s	coverage: 100.0% of statements
```
## The MiniMax algorithm

The `ComputerPlayer` struct uses the [MiniMax algorithm](https://en.wikipedia.org/wiki/Minimax) to choose its next move. Roughly speaking, the AI checks all the possible moves that could me made at each turn and chooses the one which maximises its chance of winning and assumes the opponent which minimises its chance of winning.

This is done in the following function:

```go
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
```
## Contributing

If you want to contribute to this repo please open a PR!

## License

This repo is licensed under the MIT license
