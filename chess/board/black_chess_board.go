package board

type BlackChessBoard struct {
	queenMoves []Square
	squares    BlackBoardMatrix
}

func NewBlackBoard() BlackChessBoard {
	return BlackChessBoard{
		queenMoves: []Square{},
		squares:    BlackQueenLayout,
	}
}

func (cb *BlackChessBoard) AddMove(nextMove Square) {
	cb.queenMoves = append(cb.queenMoves, nextMove)
}

func (cb *BlackChessBoard) GetQueenPosition() Square {
	return cb.queenMoves[len(cb.queenMoves)-1]
}
