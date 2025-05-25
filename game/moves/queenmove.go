package moves

type queenMove struct {
	currentSquare string
}

func newQueenMove(previous string, new string) queenMove {
	return queenMove{
		currentSquare: new,
	}
}

func (q queenMove) isFirstMove(previous *queenMove) bool {
	return previous == nil
}

// TODO: move into gamemoves
func (q queenMove) getDirection(previous *queenMove) string {
	if q.isFirstMove(previous) && q.currentSquare == "Qf8" {
		return "right"
	}

	if previous != nil && q.currentSquare > previous.currentSquare {
		return "right"
	}

	return "left"
}
