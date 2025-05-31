package board

const WHITE_QUEEN_MOVE_ZERO string = "..."

type WhiteChessBoard struct {
	queenMoves []Square
	board      WhiteBoardMatrix
}

func NewWhiteBoard() WhiteChessBoard {
	return WhiteChessBoard{
		queenMoves: []Square{},
		board:      WhiteQueenLayout,
	}
}

func (cb *WhiteChessBoard) AddMove(nextMove Square) {
	cb.queenMoves = append(cb.queenMoves, nextMove)
}

func (cb *WhiteChessBoard) GetQueenPosition() Square {
	return cb.queenMoves[len(cb.queenMoves)-1]
}

func (cb *WhiteChessBoard) Board() WhiteBoardMatrix {
	return cb.board
}

func (cb *WhiteChessBoard) GetSquareName(square Square) string {
	return cb.Board()[square.row][square.column]
}
