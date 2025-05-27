package game

import (
	"chessencryption/chess/board"
	. "chessencryption/chess/constants"
	"fmt"
)

func RunGame() {
	fmt.Print("game start: ")

	var mv *board.MoveValidator = board.NewMoveValidator()
	var blackBoard BlackChessBoard = board.NewBlackBoard(mv)
	var whiteBoard WhiteChessBoard = board.NewWhiteBoard(mv)

	blackBoard.AddMove(board.NewSquareZero("d7"))
}
