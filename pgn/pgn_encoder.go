package pgn

import (
	"fmt"
	"strings"

	"github.com/nedlir/chessencrypt/algorithm"
	"github.com/nedlir/chessencrypt/bitshandler"
	"github.com/nedlir/chessencrypt/chess/board"
)

type PGNEncoder struct {
	bitsHandler bitshandler.BitsHandler
	algo        algorithm.Algorithm
}

const WHITE_QUEEN_MOVE_ZERO = ".."
const FIRST_WHITE_SQUARE = "a6"
const FIRST_BLACK_SQUARE = "f8"
const END_OF_GAME_DRAW = "1/2-1/2"

func NewPGNEncoder() *PGNEncoder {
	return &PGNEncoder{
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
	header := fmt.Sprintf(`[Event "0000000000%d"]
[Site "?"]
[Date "????.??.??"]
[Round "?"]
[White "?"]
[Black "?"]
[Result "1/2-1/2"]
[WhiteELO "?"]
[BlackELO "?"]
[SetUp "1"]
`, chunkNumber)
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

	currentWhiteSquare := board.NewSquare(FIRST_WHITE_SQUARE, 0, 0, 0)
	currentBlackSquare := board.NewSquare(FIRST_BLACK_SQUARE, 0, 0, 5)

	squareIndex := 0
	pgnMoveIndex := 2

	for squareIndex < len(squaresToMark) {
		nextWhiteSquare, isAssist := p.algo.DetermineNextWhiteMove(currentWhiteSquare, squaresToMark[squareIndex])
		currentWhiteSquare = nextWhiteSquare

		if !isAssist {
			squareIndex++
		}

		currentBlackSquare = p.algo.DetermineNextBlackMove(isAssist, currentBlackSquare)

		pgnBoard.WriteString(fmt.Sprintf("Q%v  ", currentBlackSquare.Name()))
		pgnBoard.WriteString(fmt.Sprintf("%d. ", pgnMoveIndex))
		pgnBoard.WriteString(fmt.Sprintf("Q%v  ", nextWhiteSquare.Name()))
		pgnMoveIndex++
	}
}

func (p *PGNEncoder) writeGameResult(pgnBoard *strings.Builder) {
	pgnBoard.WriteString(END_OF_GAME_DRAW)
}
