package board

const WHITE_QUEEN_MOVE_ZERO string = "..."

type WhiteChessBoard struct {
	queenMoves []Square
	squares    WhiteBoardMatrix
}

func NewWhiteBoard(firstBitValue int) *WhiteChessBoard {

	// TODO: this will be determined by the
	if firstBitValue != 0 && firstBitValue != 1 {
		panic("first bit must be a 0 or a 1")
	}

	board := &WhiteChessBoard{
		queenMoves: []Square{},
		squares:    WhiteQueenLayout,
	}

	var zeroWhiteMove Square = NewSquare(WHITE_QUEEN_MOVE_ZERO, firstBitValue)
	board.queenMoves = append(board.queenMoves, zeroWhiteMove)

	return board
}

func (cb *WhiteChessBoard) AddMove(move Square) {
	cb.queenMoves = append(cb.queenMoves, move)
}
