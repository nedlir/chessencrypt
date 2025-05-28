package board

const WHITE_QUEEN_MOVE_ZERO string = "..."

type WhiteChessBoard struct {
	queenMoves []Square
	board      WhiteBoardMatrix
}

func NewWhiteBoard(firstBitValue int) WhiteChessBoard {
	return WhiteChessBoard{
		queenMoves: []Square{},
		board:      WhiteQueenLayout,
	}
}

func (cb *WhiteChessBoard) AddMove(move Square) {
	cb.queenMoves = append(cb.queenMoves, move)
}

func (cb *WhiteChessBoard) GetQueenPosition() Square {
	return cb.queenMoves[len(cb.queenMoves)-1]
}

func (cb *WhiteChessBoard) Board() WhiteBoardMatrix {
	return cb.board
}

func (cb *WhiteChessBoard) GetSquareName(position Position) string {
	return cb.Board()[position.row][position.column]
}
