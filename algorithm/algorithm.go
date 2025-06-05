package algorithm

import (
	"github.com/nedlir/chessencrypt/chess/board"
)

type Algorithm struct {
	movesValidator *board.MovesValidator
}

func NewAlgorithm() Algorithm {
	return Algorithm{
		movesValidator: board.NewMovesValidator(),
	}
}
