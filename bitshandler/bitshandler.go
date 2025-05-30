package bitshandler

import (
	"chessencryption/chess/board"
	"fmt"
	"math/bits"
)

type BitHandler struct {
	matrix       []byte
	currentIndex int
	allPositions []board.Position
}

func NewBitHandler(intMatrix [][]int) *BitHandler {
	byteMatrix := make([]byte, len(intMatrix))

	for row := 0; row < len(intMatrix); row++ {
		var rowByte byte = 0
		for col := 0; col < len(intMatrix[row]); col++ {
			if intMatrix[row][col] == 1 {
				rowByte |= (1 << col)
			}
		}
		byteMatrix[row] = rowByte
	}

	bh := &BitHandler{
		matrix:       byteMatrix,
		currentIndex: 0,
	}

	bh.allPositions = bh.findAllSetBits()

	return bh
}

func (bh *BitHandler) findSetBitPositions(b byte) []int {
	positions := make([]int, 0, bits.OnesCount8(b))
	for b != 0 {
		pos := bits.TrailingZeros8(b)
		positions = append(positions, pos)
		b &= b - 1
	}
	return positions
}

func (bh *BitHandler) findAllSetBits() []board.Position {
	var results []board.Position
	for rowIndex, rowByte := range bh.matrix {
		setCols := bh.findSetBitPositions(rowByte)
		for _, colIndex := range setCols {
			results = append(results, board.Position{Row: rowIndex, Col: colIndex})
		}
	}
	return results
}

func (bh *BitHandler) GetNextSetBitPosition() (board.Position, bool) {
	if bh.currentIndex >= len(bh.allPositions) {
		return board.Position{}, false // No more set bits
	}

	position := bh.allPositions[bh.currentIndex]
	bh.currentIndex++
	return position, true
}

func (bh *BitHandler) HasMoreBits() bool {
	return bh.currentIndex < len(bh.allPositions)
}

func (bh *BitHandler) GetAllSetBits() []board.Position {
	return bh.allPositions
}

func (bh *BitHandler) PrintPositions() {
	for _, pos := range bh.allPositions {
		fmt.Printf("1 at row %d, col %d\n", pos.Row(), pos.Column())
	}
}

func (bh *BitHandler) Reset() {
	bh.currentIndex = 0
}
