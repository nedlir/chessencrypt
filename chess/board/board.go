package board

import (
	. "chessencryption/chess/constants"
	"chessencryption/chess/validators"
	"strings"
)

type ChessBoard[T BoardMatrix] struct {
	color                Color
	queenMoves           []Square
	moveValidator        *validators.MoveValidator
	squares              T
	currentQueenPosition Position
}

type WhiteChessBoard = ChessBoard[WhiteBoardMatrix]
type BlackChessBoard = ChessBoard[BlackBoardMatrix]

func NewBoard(color Color, moveValidator *validators.MoveValidator) any {
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

func NewWhiteBoard(moveValidator *validators.MoveValidator) *WhiteChessBoard {
	board := &ChessBoard[WhiteBoardMatrix]{
		color:                WHITE,
		queenMoves:           []Square{WHITE_QUEEN_MOVE_ZERO},
		moveValidator:        moveValidator,
		squares:              WhiteQueenLayout,
		currentQueenPosition: WHITE_START_POSITION,
	}

	return board
}

func NewBlackBoard(moveValidator *validators.MoveValidator) *BlackChessBoard {
	board := &ChessBoard[BlackBoardMatrix]{
		color:                BLACK,
		queenMoves:           []Square{BLACK_QUEEN_STARTING_SQUARE},
		moveValidator:        moveValidator,
		squares:              BlackQueenLayout,
		currentQueenPosition: BLACK_START_POSITION,
	}

	return board
}

func (cb *ChessBoard[T]) IsNextMoveValidMove(nextMove Square) bool {
	return cb.moveValidator.IsNextMoveValidMove(cb.queenMoves, cb.color, nextMove)
}

func (cb *ChessBoard[T]) AddMove(move Square) {
	if !cb.IsNextMoveValidMove(move) {
		panic("Move inserted is invalid")
	}
	cb.queenMoves = append(cb.queenMoves, move)
}

func (cb *ChessBoard[T]) GetCurrentPosition() (Row, Column) {
	return cb.currentQueenPosition.row, cb.currentQueenPosition.column
}

func (cb *ChessBoard[T]) GetQueenPosition() Square {
	return cb.currentQueenPosition.square
}

func (cb *ChessBoard[T]) GetSquareAt(row Row, col Column) Square {
	switch squares := any(cb.squares).(type) {
	case WhiteBoardMatrix:
		if row < WhiteBoardRows && col < WhiteBoardCols {
			return squares[row][col]
		}
	case BlackBoardMatrix:
		// For 1D array, only use column index (row should be 0)
		if row == 0 && col < BlackBoardCols {
			return squares[col]
		}
	}
	return Square("")
}

func (cb *ChessBoard[T]) IsValidPosition(row Row, col Column) bool {
	switch any(cb.squares).(type) {
	case WhiteBoardMatrix:
		return row < WhiteBoardRows && col < WhiteBoardCols
	case BlackBoardMatrix:
		// For 1D array, only row 0 is valid
		return row == 0 && col < BlackBoardCols
	default:
		return false
	}
}

func (cb *ChessBoard[T]) isWhite() bool {
	return cb.color == WHITE
}

func (cb *ChessBoard[T]) isBlack() bool {
	return cb.color == BLACK
}
