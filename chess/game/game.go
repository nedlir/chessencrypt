package game

import (
	"chessencryption/bitshandler"
	"chessencryption/chess/algorithm"
	. "chessencryption/chess/board"
	"fmt"
)

func Run() {
	// for development testing, need to write tests for few cases
	// matrix without assistance moves
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 1},
	}
	// TODO: add matrix with assistance moves

	// Initialize BitHandler with the matrix
	var bitHandler *bitshandler.BitHandler = bitshandler.NewBitHandler(matrix)

	// Initialize Algorithm with BitHandler
	var algo algorithm.Algorithm = algorithm.NewAlgorithm(matrix, bitHandler)

	// var mv *board.MoveValidator = board.NewMoveValidator()
	var blackBoard BlackChessBoard = NewBlackBoard()
	var whiteBoard WhiteChessBoard = NewWhiteBoard()

	fmt.Println("My expected matrix: ")
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}

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

			isNextWhiteMoveOneStep = isReachableByWhiteQueen(nextWhiteMove.Name())

			nextBlackMove = algo.DetermineNextBlackMove(isNextWhiteMoveOneStep)

			whiteBoard.AddMove(&nextWhiteMove)
			blackBoard.AddMove(&nextBlackMove)

			algo.SetCurrentSquare(&nextWhiteMove)
		}
	}

	fmt.Println("Game finished! All set bits have been processed.")
}
