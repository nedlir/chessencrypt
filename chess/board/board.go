package board

import (
	. "chessencryption/chess/constants"
	"chessencryption/chess/validators"
	"strings"
)

type ChessBoard struct {
	color                Color
	queenMoves           []Square
	moveValidator        *validators.MoveValidator
	matrix               interface{}
	currentQueenPosition Position
}

func NewBoard(color Color, moveValidator *validators.MoveValidator) *ChessBoard {
	colorStr := Color(strings.ToLower(string(color)))
	if colorStr != WHITE && colorStr != BLACK {
		panic("invalid color: must be 'white' or 'black'")
	}

	if colorStr == WHITE {
		return NewWhiteBoard(moveValidator)
	} else {
		return NewBlackBoard(moveValidator)
	}
}

func NewWhiteBoard(moveValidator *validators.MoveValidator) *ChessBoard {
	board := &ChessBoard{
		color:                WHITE,
		queenMoves:           []Square{WHITE_QUEEN_MOVE_ZERO},
		moveValidator:        moveValidator,
		matrix:               WhiteQueenLayout,
		currentQueenPosition: WHITE_START_POSITION,
	}

	return board
}

func NewBlackBoard(moveValidator *validators.MoveValidator) *ChessBoard {
	board := &ChessBoard{
		color:                BLACK,
		queenMoves:           []Square{BLACK_QUEEN_STARTING_SQUARE},
		moveValidator:        moveValidator,
		matrix:               BlackQueenLayout,
		currentQueenPosition: BLACK_START_POSITION,
	}

	return board
}

func (cb *ChessBoard) IsNextMoveValidMove(nextMove Square) bool {
	return cb.moveValidator.IsNextMoveValidMove(cb.queenMoves, cb.color, nextMove)
}

func (cb *ChessBoard) AddMove(move Square) {
	if !cb.IsNextMoveValidMove(move) {
		panic("Move inserted is invalid")
	}
	cb.queenMoves = append(cb.queenMoves, move)
}

func (cb *ChessBoard) GetCurrentPosition() (Row, Column) {
	return cb.currentQueenPosition.row, cb.currentQueenPosition.column
}

func (cb *ChessBoard) GetQueenPosition() Square {
	return cb.currentQueenPosition.square
}

func (cb *ChessBoard) GetSquareAt(row Row, col Column) Square {
	if cb.isWhite() {
		if whiteBoard, ok := cb.matrix.([WhiteBoardRows][WhiteBoardCols]Square); ok {
			if row < WhiteBoardRows && col < WhiteBoardCols {
				return whiteBoard[row][col]
			}
		}
	} else {
		if blackBoard, ok := cb.matrix.([BlackBoardRows][BlackBoardCols]Square); ok {
			if row < BlackBoardRows && col < BlackBoardCols {
				return blackBoard[row][col]
			}
		}
	}
	return Square("")
}

func (cb *ChessBoard) IsValidPosition(row Row, col Column) bool {
	if cb.isWhite() {
		return row < WhiteBoardRows && col < WhiteBoardCols
	} else {
		return row < BlackBoardRows && col < BlackBoardCols
	}
}

func (cb *ChessBoard) isWhite() bool {
	return cb.color == WHITE
}

func (cb *ChessBoard) isBlack() bool {
	return cb.color == BLACK
}
