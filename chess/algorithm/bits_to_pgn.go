package algorithm

import (
	"fmt"

	bithandler "chessencryption/bitshandler"
	. "chessencryption/chess/board"
)

type Algorithm struct {
	bitMatrix      []byte
	currentSquare  Square
	movesValidator *MovesValidator
	bitHandler     *bithandler.BitsHandler
}

func NewAlgorithm(
	matrix []byte,
	bh *bithandler.BitsHandler,
	moves *MovesValidator,
) Algorithm {
	// start at a6 → (row=0, col=0)
	start := NewSquare("a6", 0, 0, 0)
	return Algorithm{
		bitMatrix:      matrix,
		bitHandler:     bh,
		movesValidator: moves,
		currentSquare:  start,
	}
}

func (a *Algorithm) PrintBitMatrix() {
	for row := 0; row < len(a.bitMatrix); row++ {
		fmt.Printf("Row %d (0b%08b): ", row, a.bitMatrix[row])
		for col := 0; col < 8; col++ {
			bit := (a.bitMatrix[row] >> col) & 1
			fmt.Printf("%d ", bit)
		}
		fmt.Println()
	}
}

func (a *Algorithm) DetermineNextWhiteMove(cb *WhiteChessBoard) (Square, bool) {
	targetPosition, hasNextTarget := a.bitHandler.PeekNextSetBitPosition()
	if !hasNextTarget {
		return Square{}, false //game is finished, found all necessary bits
	}

	currentPosition := a.currentSquare
	rowDifference := targetPosition.Row() - currentPosition.Row()
	columnDifference := targetPosition.Column() - currentPosition.Column()

	// 2) If it lies on the same rank/file/diagonal → consume & jump (value=1)
	if rowDifference == 0 || columnDifference == 0 || abs(rowDifference) == abs(columnDifference) {
		_, _ = a.bitHandler.FindNextSetBitPosition()
		squareName := cb.Board()[targetPosition.Row()][targetPosition.Column()]
		return NewSquare(squareName, 1, targetPosition.Row(), targetPosition.Column()), true
	}

	// 3) Otherwise: one‐step assistance → exactly one row down (value=0)
	assistanceRow := currentPosition.Row() + 1
	if assistanceRow >= WhiteBoardRowsLength {
		assistanceRow = WhiteBoardRowsLength - 1
	}
	squareName := cb.Board()[assistanceRow][currentPosition.Column()]
	return NewSquare(squareName, 0, assistanceRow, currentPosition.Column()), true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (a *Algorithm) CurrentSquare() Square {
	return a.currentSquare
}

func (a *Algorithm) SetCurrentSquare(sq *Square) {
	a.currentSquare = *sq
}

// DetermineNextBlackMove unchanged in its decision tree—only Position is removed.
func (a *Algorithm) DetermineNextBlackMove(isAssist bool, cb *WhiteChessBoard) Square {
	// map each queen‐square name to its row/col
	squarePositions := map[string]Square{
		"Qe8": NewSquare("Qe8", 0, 0, 4),
		"Qf8": NewSquare("Qf8", 0, 0, 5),
		"Qg8": NewSquare("Qg8", 0, 0, 6),
		"Qh8": NewSquare("Qh8", 0, 0, 7),
	}

	var target Square
	switch {
	case isAssist && (a.currentSquare.Name() == "Qe8" || a.currentSquare.Name() == "Qf8"):
		target = squarePositions["Qh8"]
	case isAssist && (a.currentSquare.Name() == "Qg8" || a.currentSquare.Name() == "Qh8"):
		target = squarePositions["Qe8"]
	case !isAssist && a.currentSquare.Name() == "Qe8":
		target = squarePositions["Qf8"]
	case !isAssist && a.currentSquare.Name() == "Qf8":
		target = squarePositions["Qg8"]
	case !isAssist && a.currentSquare.Name() == "Qg8":
		target = squarePositions["Qh8"]
	case !isAssist && a.currentSquare.Name() == "Qh8":
		target = squarePositions["Qg8"]
	}

	// lookup board name and assign binary value
	r, c := target.Row(), target.Column()
	boardName := cb.Board()[r][c]
	val := 0
	if !isAssist {
		val = 1
	}

	return NewSquare(boardName, val, r, c)
}
