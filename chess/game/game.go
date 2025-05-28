package game

import (
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

	var algo algorithm.Algorithm = algorithm.NewAlgorithm(matrix)
	// var mv *board.MoveValidator = board.NewMoveValidator()
	var blackBoard BlackChessBoard = NewBlackBoard()
	var whiteBoard WhiteChessBoard = NewWhiteBoard()

	fmt.Println("My expected matrix: ")
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}

	var isGameFinished bool = false
	var isNextWhiteMoveOneStep bool = false
	var nextWhiteMove Square
	var nextBlackMove Square

	for !isGameFinished {
		nextWhiteMove, isGameFinished = algo.DetermineNextWhiteMove(&whiteBoard)

		isNextWhiteMoveOneStep = isReachableByWhiteQueen(nextWhiteMove.Name())

		nextBlackMove = algo.DetermineNextBlackMove(isNextWhiteMoveOneStep)

		whiteBoard.AddMove(&nextWhiteMove)
		blackBoard.AddMove(&nextBlackMove)

		algo.SetCurrentSquare(&nextWhiteMove)
	}
}

// TODO:
// write isReachableByWhiteQueen
// print the pgn for testing
