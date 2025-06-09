package pgn

import (
	"strings"

	"github.com/nedlir/chessencrypt/algorithm"
	"github.com/nedlir/chessencrypt/chess/board"
)

type PGNDecoder struct {
	algo algorithm.Algorithm

	currentBlackMove board.Square
	currentWhiteMove board.Square
	moves            []string
	movesIndex       int
	currentByte      byte
	currentByteIndex int
	nextBlackMove    board.Square
	nextWhiteMove    board.Square
	result           []byte
	firstBitValue    int
}

func NewPGNDecoder() PGNDecoder {
	return PGNDecoder{
		algo: algorithm.NewAlgorithm(),
	}
}

func (p *PGNDecoder) PGNToBytes(pgn string) []byte {
	pgnTokens := strings.Fields(pgn)
	p.firstBitValue = determineFirstBitValue(pgnTokens)
	p.moves = extractMovesFromPgn(pgnTokens)

	p.initializeState()

	for p.movesIndex+1 < len(p.moves) {
		if p.algo.IsNewWhiteRow(p.currentWhiteMove, p.nextWhiteMove) {
			p.handleNewRow()
			continue
		}

		bit := p.calculateBitForPosition()
		p.addBitToByte(bit)
	}

	return p.finalizeResult()
}

func (p *PGNDecoder) initializeState() {
	p.currentBlackMove = board.NewSquare(FIRST_BLACK_SQUARE)
	p.currentWhiteMove = board.NewSquare(FIRST_WHITE_SQUARE)
	p.movesIndex = 0
	p.currentByte = byte(p.firstBitValue) << 7
	p.currentByteIndex = 1
	p.nextBlackMove = board.NewSquare(p.moves[p.movesIndex])
	p.nextWhiteMove = board.NewSquare(p.moves[p.movesIndex+1])
	p.result = []byte{}
}

func (p *PGNDecoder) handleNewRow() {
	if p.currentByteIndex > 0 {
		p.result = append(p.result, p.currentByte)
	}
	p.currentByte = 0
	p.currentByteIndex = 0

	p.currentWhiteMove = p.nextWhiteMove
	p.currentBlackMove = p.nextBlackMove
	p.movesIndex += 2
	if p.movesIndex+1 < len(p.moves) {
		p.nextBlackMove = board.NewSquare(p.moves[p.movesIndex])
		p.nextWhiteMove = board.NewSquare(p.moves[p.movesIndex+1])
	}
}

func (p *PGNDecoder) calculateBitForPosition() byte {
	var bit byte
	if p.currentByteIndex == p.nextWhiteMove.Column() {
		if p.algo.IsAssistanceMove(p.currentBlackMove, p.nextBlackMove) {
			bit = 0
		} else {
			bit = 1
		}
		p.currentBlackMove = p.nextBlackMove
		p.currentWhiteMove = p.nextWhiteMove
		p.movesIndex += 2
		if p.movesIndex+1 < len(p.moves) {
			p.nextBlackMove = board.NewSquare(p.moves[p.movesIndex])
			p.nextWhiteMove = board.NewSquare(p.moves[p.movesIndex+1])
		}
	} else {
		bit = 0
	}
	return bit
}

func (p *PGNDecoder) addBitToByte(bit byte) {
	p.currentByte |= bit << (7 - p.currentByteIndex)
	if p.currentByteIndex == 7 {
		p.result = append(p.result, p.currentByte)
		p.currentByte = 0
		p.currentByteIndex = 0
	} else {
		p.currentByteIndex++
	}
}

func (p *PGNDecoder) finalizeResult() []byte {
	if p.currentByteIndex > 0 {
		p.result = append(p.result, p.currentByte)
	}
	return p.result
}

func (p *PGNDecoder) Reset() {
	p.currentBlackMove = board.Square{}
	p.currentWhiteMove = board.Square{}
	p.moves = nil
	p.movesIndex = 0
	p.currentByte = 0
	p.currentByteIndex = 0
	p.nextBlackMove = board.Square{}
	p.nextWhiteMove = board.Square{}
	p.result = nil
	p.firstBitValue = 0
}

func determineFirstBitValue(pgnTokens []string) int {
	firstFenValue := extractFirstFenValue(pgnTokens)
	if firstFenValue == board.WhiteKingOnA8 {
		return 0
	} else if firstFenValue == board.WhiteKingOnA7 {
		return 1
	}
	return -1
}

func extractFirstFenValue(pgnTokens []string) string {
	rawFen := pgnTokens[21]
	return rawFen[1:2]
}

func extractMovesFromPgn(pgnTokens []string) []string {
	gameMoveTokens := pgnTokens[28:]

	moves := make([]string, 0)
	for i, token := range gameMoveTokens {
		if i%3 != 1 { // valid move index
			// Strip 'Q' prefix if present
			if len(token) > 0 && token[0] == 'Q' {
				token = token[1:]
			}
			moves = append(moves, token)
		}
	}

	return moves
}
