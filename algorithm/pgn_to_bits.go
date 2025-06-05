package algorithm

import "github.com/nedlir/chessencrypt/chess/board"

// func DetermineNextSetBit() {
// 	nextBitToSet
// }

func calculatePositionalDistance(currentSquare string, nextSquare string) int {
	currentColumn := board.BlackSquarePositions[currentSquare].Col
	nextColumn := board.BlackSquarePositions[nextSquare].Col

	return currentColumn - nextColumn
}
