package game

import (
	"chessencryption/bitshandler"
	"chessencryption/chess/algorithm"
	. "chessencryption/chess/board"
	"fmt"
)

func Run() {
	// Matrix represented as bytes for efficient bit operations
	// TODO: produce more matrices like this with test cases for edge cases and different values
	matrixRows := []byte{
		0b01011010, // row 0: bits at positions 1,3,4,6
		0b10100101, // row 1: bits at positions 0,2,5,7
		0b00110110, // row 2: bits at positions 1,2,5,6
		0b11001001, // row 3: bits at positions 0,3,6,7
		0b01101100, // row 4: bits at positions 2,3,5,6
		0b10010011, // row 5: bits at positions 0,1,4,7
	}

	var bitHandler *bitshandler.BitHandler = bitshandler.NewBitHandler(matrixRows)

	var moveValidator *MoveValidator = NewMoveValidator()

	var algo algorithm.Algorithm = algorithm.NewAlgorithm(matrixRows, bitHandler, moveValidator)

	var blackBoard BlackChessBoard = NewBlackBoard()
	var whiteBoard WhiteChessBoard = NewWhiteBoard()

	fmt.Println(" expected matrix (as bytes):")
	algo.PrintBitMatrix()

	fmt.Println("\nSet bit positions found by BitHandler:")
	bitHandler.PrintPositions()
	fmt.Println()

	var isGameFinished bool = false
	var isNextWhiteMoveOneStep bool = false
	var nextWhiteMove Square
	var nextBlackMove Square

	for !isGameFinished {
		nextWhiteMove, isGameFinished = algo.DetermineNextWhiteMove(&whiteBoard)

		if !isGameFinished {
			fmt.Printf("Next white move: %s at position (%d, %d)\n",
				nextWhiteMove.Name(),
				nextWhiteMove.Position().Row(),
				nextWhiteMove.Position().Column())

			isNextWhiteMoveOneStep = moveValidator.IsNextMoveValidMove(algo.CurrentSquare(), nextWhiteMove)

			nextBlackMove = algo.DetermineNextBlackMove(isNextWhiteMoveOneStep, &whiteBoard)

			whiteBoard.AddMove(&nextWhiteMove)
			blackBoard.AddMove(&nextBlackMove)

			algo.SetCurrentSquare(&nextWhiteMove)
		}
	}

	fmt.Println("Game finished! All set bits have been processed.")
}
