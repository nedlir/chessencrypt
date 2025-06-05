package bitshandler

import (
	"math/bits"
	"slices"

	"github.com/nedlir/chessencrypt/chess/board"
)

type BitsHandler struct {
	matrix       []byte
	currentIndex int
	allPositions []board.Square
}

func NewBitHandler() BitsHandler {
	return BitsHandler{}
}

func (bh *BitsHandler) UpdateMatrix(newMatrix []byte) {
	bh.matrix = newMatrix
	bh.allPositions = bh.findAllSetBits()
	bh.currentIndex = 0 // Reset iterator
}

func (bh *BitsHandler) findAllSetBits() []board.Square {
	var squares []board.Square

	for rowIndex, rowByte := range bh.matrix {
		if rowByte == 0 {
			continue // Skip empty rows (0b00000000) entirely
		}

		rowStartIndex := len(squares)

		currentByte := rowByte
		for currentByte != 0 {
			squares = bh.addSquare(squares, rowIndex, currentByte)
			currentByte = removeRightmostSetBit(currentByte)
		}

		if len(squares) > rowStartIndex {
			slices.Reverse(squares[rowStartIndex:])
		}
	}

	return squares
}

func (bh *BitsHandler) addSquare(squares []board.Square, rowIndex int, currentByte byte) []board.Square {
	nextSetBitPosition := getRightmostSetBitLeftToRightPos(currentByte)
	squareName := board.WhiteQueenLayout[rowIndex][nextSetBitPosition]
	return append(squares, board.NewSquare(squareName))
}
func removeRightmostSetBit(currentByte byte) byte {
	return currentByte & (currentByte - 1)
}

func getRightmostSetBitLeftToRightPos(currentByte byte) int {
	rightToLeftPos := bits.TrailingZeros8(currentByte)
	leftToRightPos := 7 - rightToLeftPos
	return leftToRightPos
}

func (bh *BitsHandler) AllSetBits() []board.Square {
	return bh.allPositions
}

func (bh *BitsHandler) Reset() {
	bh.currentIndex = 0
}

func (bh *BitsHandler) CurrentIndex() int {
	return bh.currentIndex
}

func (bh *BitsHandler) IsFirstBitZero() bool {
	return (bh.matrix[0]>>7)&1 == 0
}
