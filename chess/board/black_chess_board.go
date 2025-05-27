package board

type BlackChessBoard struct {
	queenMoves []Square
	squares    BlackBoardMatrix
}

func NewBlackBoard() *BlackChessBoard {
	board := &BlackChessBoard{
		queenMoves: []Square{},
		squares:    BlackQueenLayout,
	}
	return board
}

// TODO: move logic to algorithm.go
func (cb *BlackChessBoard) AddMove(isNextMoveAssistance bool) {
	currentPosition := cb.GetQueenPosition().Name()
	var nextMove string
	if isNextMoveAssistance {
		switch currentPosition {
		case "Qe8", "Qf8":
			nextMove = "Qh8"
		case "Qg8", "Qh8":
			nextMove = "Qe8"
		}
	} else { // next move is supposed to mark the bit
		switch currentPosition {
		case "Qe8":
			nextMove = "Qf8"
		case "Qf8":
			nextMove = "Qg8"
		case "Qg8":
			nextMove = "Qh8"
		case "Qh8":
			nextMove = "Qg8"
		}

	}

	cb.queenMoves = append(cb.queenMoves, NewSquare(nextMove, 0))
}

func (cb *BlackChessBoard) GetQueenPosition() Square {
	return cb.queenMoves[len(cb.queenMoves)-1]
}
