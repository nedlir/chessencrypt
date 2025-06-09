package pgn

import (
	"fmt"
	"strings"

	"github.com/nedlir/chessencrypt/algorithm"
	"github.com/nedlir/chessencrypt/chess/board"
	"github.com/nedlir/chessencrypt/utils/bitshandler"
)

type PGNEncoder struct {
	bitsHandler        bitshandler.BitsHandler
	algo               algorithm.Algorithm
	currentWhiteSquare board.Square
	currentBlackSquare board.Square
	squareIndex        int
	squaresToMark      []board.Square
	pgnMoveIndex       int
}

func NewPGNEncoder() PGNEncoder {
	return PGNEncoder{
		bitsHandler: bitshandler.NewBitHandler(),
		algo:        algorithm.NewAlgorithm(),
	}
}

func (p *PGNEncoder) BytesToPgn(matrix []byte, chunkNumber int) string {
	p.bitsHandler.UpdateMatrix(matrix)
	p.initializeState()

	var pgnBoard strings.Builder
	pgnBoard.WriteString(p.buildHeader(chunkNumber))

	for p.squareIndex < len(p.squaresToMark) {
		whiteMove, blackMove, isAssist := p.generateNextMove()
		p.writeMove(&pgnBoard, whiteMove, blackMove)
		p.advanceState(isAssist)
	}

	p.writeGameResult(&pgnBoard)
	return pgnBoard.String()
}

func (p *PGNEncoder) initializeState() {
	p.squaresToMark = p.bitsHandler.AllSetBits()
	p.currentWhiteSquare = board.NewSquare(FIRST_WHITE_SQUARE)
	p.currentBlackSquare = board.NewSquare(FIRST_BLACK_SQUARE)
	p.squareIndex = 0
	p.pgnMoveIndex = 2

	if len(p.squaresToMark) > 0 && p.squaresToMark[p.squareIndex].Name() == FIRST_WHITE_SQUARE {
		p.squareIndex++
	}
}

func (p *PGNEncoder) buildHeader(chunkNumber int) string {
	var header strings.Builder

	header.WriteString(fmt.Sprintf(PGN_HEADER_TEMPLATE, chunkNumber))
	header.WriteString(p.buildFENString())
	header.WriteString(fmt.Sprintf("1.%s ", WHITE_QUEEN_MOVE_ZERO))

	return header.String()
}

func (p *PGNEncoder) buildFENString() string {
	if p.bitsHandler.IsFirstBitZero() {
		return fmt.Sprintf(`[FEN "%v"]`+"\n", board.FENZero)
	}
	return fmt.Sprintf(`[FEN "%v"]`+"\n", board.FENOne)
}

func (p *PGNEncoder) generateNextMove() (whiteMove, blackMove string, isAssist bool) {
	targetSquare := p.squaresToMark[p.squareIndex]

	nextWhiteSquare, isAssist := p.algo.DetermineNextWhiteMove(p.currentWhiteSquare, targetSquare)
	p.currentWhiteSquare = nextWhiteSquare

	nextBlackSquare := p.algo.DetermineNextBlackMove(isAssist, p.currentBlackSquare)
	p.currentBlackSquare = nextBlackSquare

	whiteMove = p.formatMove(nextWhiteSquare)
	blackMove = p.formatMove(nextBlackSquare)

	return whiteMove, blackMove, isAssist
}

func (p *PGNEncoder) formatMove(square board.Square) string {
	return fmt.Sprintf("Q%v", square.Name())
}

func (p *PGNEncoder) writeMove(pgnBoard *strings.Builder, whiteMove, blackMove string) {
	pgnBoard.WriteString(fmt.Sprintf("%s  ", blackMove))
	pgnBoard.WriteString(fmt.Sprintf("%d. ", p.pgnMoveIndex))
	pgnBoard.WriteString(fmt.Sprintf("%s  ", whiteMove))
}

func (p *PGNEncoder) advanceState(isAssist bool) {
	if !isAssist {
		p.squareIndex++
	}
	p.pgnMoveIndex++
}

func (p *PGNEncoder) writeGameResult(pgnBoard *strings.Builder) {
	pgnBoard.WriteString(END_OF_GAME_DRAW)
}

func (p *PGNEncoder) Reset() {
	p.currentWhiteSquare = board.Square{}
	p.currentBlackSquare = board.Square{}
	p.squareIndex = 0
	p.squaresToMark = nil
	p.pgnMoveIndex = 0
}
