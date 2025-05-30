package algorithm

import (
	bithandler "chessencryption/bitshandler"
	. "chessencryption/chess/board"
	. "chessencryption/chess/fen"
	"fmt"
)

type Algorithm struct {
	bitMatrix     []byte
	currentSquare Square
	moveValidator *MoveValidator
	bitHandler    *bithandler.BitsHandler
}

func NewAlgorithm(b []byte, bh *bithandler.BitsHandler, mv *MoveValidator) Algorithm {
	return Algorithm{
		bitMatrix:     b,
		bitHandler:    bh,
		moveValidator: mv,
		currentSquare: NewSquare(
			"a6",
			0,
			NewPosition(0, 1),
		),
	}
}

func (a *Algorithm) PrintBitMatrix() {
	fmt.Println("Bit matrix representation:")
	for row := 0; row < len(a.bitMatrix); row++ {
		fmt.Printf("Row %d (0b%08b): ", row, a.bitMatrix[row])
		for col := 0; col < 8; col++ {
			bit := (a.bitMatrix[row] >> col) & 1
			fmt.Printf("%d ", bit)
		}
		fmt.Println()
	}
}

func (a *Algorithm) DetermineFEN(fen string) string {
	var firstBit int = int(a.bitMatrix[0] & 1)

	if firstBit == 0 {
		return FENZero
	} else {
		return FENOne
	}
}

func (a *Algorithm) DetermineNextWhiteMove(cb *WhiteChessBoard) (square Square, isExist bool) {
	position, found := a.bitHandler.FindNextSetBitPosition()

	if !found {
		return Square{}, false
	}

	if position.Row() < WhiteBoardRowsLength && position.Column() < WhiteBoardColsLength {
		squareName := cb.Board()[position.Row()][position.Column()]
		newSquare := NewSquare(
			squareName,
			1,
			NewPosition(position.Row(), position.Column()),
		)

		if a.moveValidator != nil && a.currentSquare.Name() != "" {
			valid := a.moveValidator.IsNextMoveValidMove(a.currentSquare, newSquare)
			if !valid {
				fmt.Printf("Invalid move from %s to %s, skipping...\n", a.currentSquare.Name(), squareName)
			}
		}

		return newSquare, true
	}

	return Square{}, false
}

func (a *Algorithm) CurrentSquare() Square {
	return a.currentSquare
}

func (a *Algorithm) SetCurrentSquare(square *Square) {
	a.currentSquare = *square
}

func (a *Algorithm) DetermineNextBlackMove(isNextMoveAssistance bool, cb *WhiteChessBoard) Square {
	var nextSquare Square
	var targetPosition Position

	squarePositions := map[string]Position{
		"Qe8": NewPosition(0, 4),
		"Qf8": NewPosition(0, 5),
		"Qg8": NewPosition(0, 6),
		"Qh8": NewPosition(0, 7),
	}

	if isNextMoveAssistance {
		switch a.currentSquare.Name() {
		case "Qe8", "Qf8":
			targetPosition = squarePositions["Qh8"]
		case "Qg8", "Qh8":
			targetPosition = squarePositions["Qe8"]
		}
	} else { // next move is supposed to mark the bit
		switch a.currentSquare.Name() {
		case "Qe8":
			targetPosition = squarePositions["Qf8"]
		case "Qf8":
			targetPosition = squarePositions["Qg8"]
		case "Qg8":
			targetPosition = squarePositions["Qh8"]
		case "Qh8":
			targetPosition = squarePositions["Qg8"]
		}
	}

	squareName := GetSquareName(targetPosition, cb.Board())

	binaryValue := 0
	if !isNextMoveAssistance {
		binaryValue = 1
	}

	nextSquare = NewSquare(squareName, binaryValue, targetPosition)
	return nextSquare
}

func (a *Algorithm) GetBitHandler() *bithandler.BitsHandler {
	return a.bitHandler
}

func (a *Algorithm) GetMoveValidator() *MoveValidator {
	return a.moveValidator
}
