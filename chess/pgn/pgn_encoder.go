package pgn

import (
	"fmt"
	"strings"

	"github.com/nedlir/chessencrypt/algorithm"
	"github.com/nedlir/chessencrypt/chess/board"
	"github.com/nedlir/chessencrypt/utils/bitshandler"
)

type PGNEncoder struct {
	bitsHandler bitshandler.BitsHandler
	algo        algorithm.Algorithm
}

func NewPGNEncoder() PGNEncoder {
	return PGNEncoder{
		bitsHandler: bitshandler.NewBitHandler(),
		algo:        algorithm.NewAlgorithm(),
	}
}

func (p *PGNEncoder) BytesToPgn(matrix []byte, chunkNumber int) string {
	p.bitsHandler.UpdateMatrix(matrix)

	var pgnBoard strings.Builder

	p.writePGNHeaders(&pgnBoard, chunkNumber)

	p.writeMoves(&pgnBoard)

	p.writeGameResult(&pgnBoard)

	return pgnBoard.String()
}

func (p *PGNEncoder) writePGNHeaders(pgnBoard *strings.Builder, chunkNumber int) {

	header := fmt.Sprintf(PGN_HEADER_TEMPLATE, chunkNumber)
	pgnBoard.WriteString(header)

	if p.bitsHandler.IsFirstBitZero() {
		pgnBoard.WriteString(fmt.Sprintf(`[FEN "%v"]`+"\n", board.FENZero))
	} else {
		pgnBoard.WriteString(fmt.Sprintf(`[FEN "%v"]`+"\n", board.FENOne))
	}

	pgnBoard.WriteString(fmt.Sprintf("1.%s ", WHITE_QUEEN_MOVE_ZERO))
}

func (p *PGNEncoder) writeMoves(pgnBoard *strings.Builder) {
	squaresToMark := p.bitsHandler.AllSetBits()

	currentWhiteSquare := board.NewSquare(FIRST_WHITE_SQUARE)
	currentBlackSquare := board.NewSquare(FIRST_BLACK_SQUARE)

	squareIndex := 0
	if len(squaresToMark) > 0 && squaresToMark[squareIndex].Name() == FIRST_WHITE_SQUARE {
		squareIndex++
	}

	pgnMoveIndex := 2

	for squareIndex < len(squaresToMark) {

		targetSquare := squaresToMark[squareIndex]

		nextWhiteSquare, isAssist := p.algo.DetermineNextWhiteMove(currentWhiteSquare, targetSquare)

		currentWhiteSquare = nextWhiteSquare

		if !isAssist {
			squareIndex++
		}

		currentBlackSquare = p.algo.DetermineNextBlackMove(isAssist, currentBlackSquare)

		blackMove := fmt.Sprintf("Q%v  ", currentBlackSquare.Name())
		moveNumber := fmt.Sprintf("%d. ", pgnMoveIndex)
		whiteMove := fmt.Sprintf("Q%v  ", nextWhiteSquare.Name())

		pgnBoard.WriteString(blackMove)
		pgnBoard.WriteString(moveNumber)
		pgnBoard.WriteString(whiteMove)

		pgnMoveIndex++
	}

}

func (p *PGNEncoder) writeGameResult(pgnBoard *strings.Builder) {
	pgnBoard.WriteString(END_OF_GAME_DRAW)
}
