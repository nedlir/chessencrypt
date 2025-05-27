package board

const WHITE_QUEEN_MOVE_ZERO string = "..."

type WhiteChessBoard struct {
	queenMoves    []Square
	moveValidator *MoveValidator
	squares       WhiteBoardMatrix
}

func NewWhiteBoard(mv *MoveValidator, firstBitValue int) *WhiteChessBoard {

	// TODO: this will be determined by the
	if firstBitValue != 0 && firstBitValue != 1 {
		panic("first bit must be a 0 or a 1")
	}

	board := &WhiteChessBoard{
		queenMoves:    []Square{},
		moveValidator: mv,
		squares:       WhiteQueenLayout,
	}

	var zeroWhiteMove Square = NewSquare(WHITE_QUEEN_MOVE_ZERO, firstBitValue)
	board.queenMoves = append(board.queenMoves, zeroWhiteMove)

	return board
}

func (cb *WhiteChessBoard) IsNextMoveValidMove(nextMove Square) bool {
	return cb.moveValidator.IsNextMoveValidMove(cb.queenMoves, nextMove)
}

func (cb *WhiteChessBoard) AddMove(move Square) {
	if !cb.IsNextMoveValidMove(move) {
		panic("Move inserted is invalid")
	}
	cb.queenMoves = append(cb.queenMoves, move)
}
