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

func (bh *BitsHandler) findSetBitPositions(b byte) []int {
	positions := make([]int, 0, bits.OnesCount8(b))
	for b != 0 {
		pos := bits.TrailingZeros8(b)
		positions = append(positions, pos)
		b &= b - 1
	}
	return positions
}

func (bh *BitsHandler) findAllSetBits() []board.Square {
	var res []board.Square
	for r, row := range bh.matrix {
		for _, c := range bh.findSetBitPositions(row) {
			res = append(res, board.NewSquare("", 0, r, c))
		}
	}
	return res
}

// PeekNextSetBitPosition returns the next set-bit WITHOUT advancing currentIndex
func (bh *BitsHandler) PeekNextSetBitPosition() (board.Square, bool) {
	if bh.currentIndex >= len(bh.allPositions) {
		return board.Square{}, false
	}
	return bh.allPositions[bh.currentIndex], true
}

// FindNextSetBitPosition returns the next set-bit AND advances currentIndex
func (bh *BitsHandler) FindNextSetBitPosition() (board.Square, bool) {
	if bh.currentIndex >= len(bh.allPositions) {
		return board.Square{}, false
	}
	pos := bh.allPositions[bh.currentIndex]
	bh.currentIndex++
	return pos, true
}

func (bh *BitsHandler) HasMoreBits() bool {
	return bh.currentIndex < len(bh.allPositions)
}

func (bh *BitsHandler) AllSetBits() []board.Square {
	return bh.allPositions
}

func (bh *BitsHandler) PrintPositions() {
	for _, pos := range bh.allPositions {
		fmt.Printf("1 at row %d, col %d\n", pos.Row(), pos.Column())
	}
}

func (bh *BitsHandler) Reset() {
	bh.currentIndex = 0
}

func (bh *BitsHandler) Matrix() []byte {
	return bh.matrix
}
