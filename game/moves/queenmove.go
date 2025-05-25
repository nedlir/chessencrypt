package moves

type queenMove struct {
	square square
}

func newQueenMove(newSquare square) queenMove {
	return queenMove{
		square: newSquare,
	}
}

func (qm *queenMove) Square() square {
	return qm.square
}
