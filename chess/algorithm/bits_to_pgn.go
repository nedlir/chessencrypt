package algorithm

import (
	"chessencryption/chess/board"
)

type Algorithm struct {
	bitMatrix      []byte
	movesValidator *board.MovesValidator
}

func NewAlgorithm(
	matrix []byte,
) Algorithm {
	return Algorithm{
		bitMatrix:      matrix,
		movesValidator: board.NewMovesValidator(),
	}
}

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
	binaryValue := 0
	name := board.WhiteQueenLayout[row][col]

	return board.NewSquare(name, binaryValue, row, col)
}

func (a *Algorithm) DetermineNextBlackMove(isAssist bool, currentSquare board.Square) board.Square {
	if isAssist {
		switch currentSquare.Name() {
		case "e8", "f8":
			return board.NewSquare("h8", 0, 0, 7)
		case "g8", "h8":
			return board.NewSquare("e8", 0, 0, 4)
		}
	} else {
		switch currentSquare.Name() {
		case "e8":
			return board.NewSquare("f8", 0, 0, 5)
		case "f8":
			return board.NewSquare("g8", 0, 0, 6)
		case "g8":
			return board.NewSquare("h8", 0, 0, 7)
		case "h8":
			return board.NewSquare("g8", 0, 0, 6)
		}
	}

	return board.Square{}
}
