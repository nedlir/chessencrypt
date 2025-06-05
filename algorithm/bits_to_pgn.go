package algorithm

import (
	"github.com/nedlir/chessencrypt/chess/board"
)

func (a *Algorithm) DetermineNextWhiteMove(currentSquare board.Square, nextSquareWithOne board.Square) (board.Square, bool) {
	if a.movesValidator.IsNextMoveValidMove(currentSquare, nextSquareWithOne) {
		return nextSquareWithOne, false
	}

	traversalDownSquare := getTraversalDownSquare(currentSquare, nextSquareWithOne)

	return traversalDownSquare, true
}

func getTraversalDownSquare(currentSquare board.Square, nextSquareWithOne board.Square) board.Square {
	row := nextSquareWithOne.Row()
	col := currentSquare.Column()
	name := board.WhiteQueenLayout[row][col]

	return board.NewSquare(name)
}

func (a *Algorithm) DetermineNextBlackMove(isAssist bool, currentSquare board.Square) board.Square {
	if isAssist {
		switch currentSquare.Name() {
		case "e8", "f8":
			return board.NewSquare("h8")
		case "g8", "h8":
			return board.NewSquare("e8")
		}
	} else {
		switch currentSquare.Name() {
		case "e8":
			return board.NewSquare("f8")
		case "f8":
			return board.NewSquare("g8")
		case "g8":
			return board.NewSquare("h8")
		case "h8":
			return board.NewSquare("g8")
		}
	}

	return board.Square{}
}
