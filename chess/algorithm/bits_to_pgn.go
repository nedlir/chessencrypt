package algorithm

import (
	"fmt"

	bithandler "chessencryption/bitshandler"
	. "chessencryption/chess/board"
)

type Algorithm struct {
	bitMatrix     []byte
	currentSquare Square
	moveValidator *MoveValidator
	bitHandler    *bithandler.BitsHandler
}

func NewAlgorithm(
	matrix []byte,
	bh *bithandler.BitsHandler,
	mv *MoveValidator,
) Algorithm {
	// start at a6 → (row=0, col=0)
	start := NewSquare("a6", 0, 0, 0)
	return Algorithm{
		bitMatrix:     matrix,
		bitHandler:    bh,
		moveValidator: mv,
		currentSquare: start,
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

// DetermineNextWhiteMove:
//   - Peek next set‐bit
//   - If directly reachable, consume it & return value=1
//   - Else step one row down (row+1) with value=0
func (a *Algorithm) DetermineNextWhiteMove(cb *WhiteChessBoard) (Square, bool) {
	// 1) Peek the next 1‐bit, but do NOT advance the pointer yet:
	targetPos, ok := a.bitHandler.PeekNextSetBitPosition()
	if !ok {
		return Square{}, false
	}

	curr := a.currentSquare
	dr := targetPos.Row() - curr.Row()
	dc := targetPos.Column() - curr.Column()

	// 2) If it lies on the same rank/file/diagonal → consume & jump (value=1)
	if dr == 0 || dc == 0 || abs(dr) == abs(dc) {
		_, _ = a.bitHandler.FindNextSetBitPosition()
		name := cb.Board()[targetPos.Row()][targetPos.Column()]
		return NewSquare(name, 1, targetPos.Row(), targetPos.Column()), true
	}

	// 3) Otherwise: one‐step assistance → exactly one row down (value=0)
	nextRow := curr.Row() + 1
	if nextRow >= WhiteBoardRowsLength {
		nextRow = WhiteBoardRowsLength - 1
	}
	name := cb.Board()[nextRow][curr.Column()]
	return NewSquare(name, 0, nextRow, curr.Column()), true
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
