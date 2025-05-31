package bitshandler

import (
	"chessencryption/chess/board"
	"fmt"
	"math/bits"
)

type BitsHandler struct {
	matrix       []byte
	currentIndex int
	allPositions []board.Square
}

func NewBitHandler(byteMatrix []byte) *BitsHandler {
	bh := &BitsHandler{matrix: byteMatrix}
	bh.allPositions = bh.findAllSetBits()
	return bh
}

// / findSetBitPositions returns the bit‐offsets of all 1s in a single row byte.
// Uses left-to-right indexing (MSB=0 to LSB=7) and returns in left-to-right order
func (bh *BitsHandler) findSetBitPositions(rowByte byte) []int {
	// Preallocate for the number of 1‐bits in this byte
	positions := make([]int, 0, bits.OnesCount8(rowByte))

	for rowByte != 0 {
		// Find rightmost set bit using TrailingZeros8
		rightToLeftPos := bits.TrailingZeros8(rowByte)

		// Convert to left-to-right indexing (MSB=0 to LSB=7)
		leftToRightPos := 7 - rightToLeftPos

		positions = append(positions, leftToRightPos)

		// Clear the lowest set bit
		rowByte &= rowByte - 1
	}

	// Reverse the slice to get left-to-right order
	for i, j := 0, len(positions)-1; i < j; i, j = i+1, j-1 {
		positions[i], positions[j] = positions[j], positions[i]
	}

	return positions
}

// findAllSetBits walks the entire byte‐matrix, finds every bit set to 1,
// and turns each into a board.Square(rowIndex, colIndex).
func (bh *BitsHandler) findAllSetBits() []board.Square {
	var squares []board.Square

	for rowIndex, rowByte := range bh.matrix {
		// find all column‐indices with a 1‐bit (using left-to-right indexing)
		setCols := bh.findSetBitPositions(rowByte)

		for _, colIndex := range setCols {
			// create a Square at (rowIndex, colIndex)
			squares = append(squares, board.NewSquare("", 0, rowIndex, colIndex))
		}
	}

	return squares
}

func (bh *BitsHandler) PeekNextSetBitPosition() (board.Square, bool) {
	if bh.currentIndex >= len(bh.allPositions) {
		return board.Square{}, false
	}
	return bh.allPositions[bh.currentIndex], true
}

func (bh *BitsHandler) FindNextSetBitPosition() (board.Square, bool) {
	if bh.currentIndex >= len(bh.allPositions) {
		return board.Square{}, false
	}
	sq := bh.allPositions[bh.currentIndex]
	bh.currentIndex++
	return sq, true
}

func (bh *BitsHandler) HasMoreBits() bool {
	return bh.currentIndex < len(bh.allPositions)
}

func (bh *BitsHandler) AllSetBits() []board.Square {
	return bh.allPositions
}

func (bh *BitsHandler) PrintPositions() {
	for _, pos := range bh.allPositions {
		fmt.Printf("(r: %d, col: %d)\n", pos.Row(), pos.Column())
	}
}

func (bh *BitsHandler) Reset() {
	bh.currentIndex = 0
}

func (bh *BitsHandler) Matrix() []byte {
	return bh.matrix
}
