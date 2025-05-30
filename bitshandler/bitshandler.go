package bithandler

import (
	"chessencryption/chess/board"
	"fmt"
	"math/bits"
)

type BitsHandler struct {
	matrix       []byte
	currentIndex int
	allPositions []board.Position
}

func NewBitHandler(byteMatrix []byte) *BitsHandler {
	bh := &BitsHandler{
		matrix:       byteMatrix,
		currentIndex: 0,
	}

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

func (bh *BitsHandler) findAllSetBits() []board.Position {
	var results []board.Position
	for rowIndex, rowByte := range bh.matrix {
		setCols := bh.findSetBitPositions(rowByte)
		for _, colIndex := range setCols {
			results = append(results, board.NewPosition(rowIndex, colIndex))
		}
	}
	return results
}

func (bh *BitsHandler) FindNextSetBitPosition() (board.Position, bool) {
	if bh.currentIndex >= len(bh.allPositions) {
		return board.Position{}, false // No more set bits
	}

	position := bh.allPositions[bh.currentIndex]
	bh.currentIndex++
	return position, true
}

func (bh *BitsHandler) HasMoreBits() bool {
	return bh.currentIndex < len(bh.allPositions)
}

func (bh *BitsHandler) AllSetBits() []board.Position {
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
