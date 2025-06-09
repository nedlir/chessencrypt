package algorithm

import (
	"math"

	"github.com/nedlir/chessencrypt/chess/board"
)

func (a *Algorithm) IsNewWhiteRow(currentWhiteMove board.Square, nextWhiteMove board.Square) bool {
	return math.Abs(float64(currentWhiteMove.Row()-nextWhiteMove.Row())) >= 1
}

func (a *Algorithm) IsAssistanceMove(currentBlackMove board.Square, nextBlackMove board.Square) bool {
	return math.Abs(float64(currentBlackMove.Column()-nextBlackMove.Column())) > 1
}
