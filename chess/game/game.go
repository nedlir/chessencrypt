package game

import (
	"chessencryption/chess/algorithm"
	"fmt"
)

func Run() {
	// for development testing, need to write tests for few cases
	// matrix without assistance moves
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1},
	}
	// TODO: add matrix with assistance moves

	var algo algorithm.Algorithm = algorithm.NewAlgorithm(matrix)
	// var mv *board.MoveValidator = board.NewMoveValidator()
	// var blackBoard *board.BlackChessBoard = board.NewBlackBoard()
	// var whiteBoard *board.WhiteChessBoard = board.NewWhiteBoard(0)

	fmt.Println("My expected matrix: ")
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}

	for !algo.IsGameFinished() {
		algo.FindNextBitToSet()

		target := algo.NextPosition()
		fmt.Printf("Next 1st bit square: [%d][%d]\n", target.Row(), target.Column())

		// right now the queen hasn't moved yet
		cur := algo.CurrentPosition()
		fmt.Printf("currently at:       [%d][%d]\n\n", cur.Row(), cur.Column())

		algo.ApplyNextMove()
	}

	// blackBoard.AddMove(board.NewSquareZero("d7"))
}
