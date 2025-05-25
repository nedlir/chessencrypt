package moves

type QueenMove struct {
	currentSquare string
}

func NewQueenMove(current string) QueenMove {
	return QueenMove{
		currentSquare: current,
	}
}

func (q QueenMove) isFirstMove(previous *QueenMove) bool {
	return previous == nil
}

func (q QueenMove) GetDirection(previous *QueenMove) string {
	if q.isFirstMove(previous) && q.currentSquare == "Qf8" {
		return "right"
	}

	if previous != nil && q.currentSquare > previous.currentSquare {
		return "right"
	}

	return "left"
}
