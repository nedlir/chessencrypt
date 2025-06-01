package algorithm

import (
	"chessencryption/chess/board"
	"fmt"
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

func (a *Algorithm) PrintBitMatrix() {
	for row := 0; row < len(a.bitMatrix); row++ {
		fmt.Printf("Row %d (0b%08b): ", row, a.bitMatrix[row])
		for col := 0; col < 8; col++ {
			bit := (a.bitMatrix[row] >> (7 - col)) & 1 //  start from MSB
			fmt.Printf("%d ", bit)
		}
		fmt.Println()
	}
}

func (a *Algorithm) DetermineNextWhiteMove(currentSquare board.Square, nextTargetSquareWithOne board.Square) (board.Square, bool) {
	if a.movesValidator.IsNextMoveValidMove(currentSquare, nextTargetSquareWithOne) {
		return nextTargetSquareWithOne, false
	}

	fmt.Println("traversalDownSquare:")
	traversalDownSquare := getTraversalDownSquare(currentSquare, nextTargetSquareWithOne)

	return traversalDownSquare, true

}

func getTraversalDownSquare(currentSquare board.Square, nextSquare board.Square) board.Square {
	row := nextSquare.Row() - currentSquare.Row()
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
