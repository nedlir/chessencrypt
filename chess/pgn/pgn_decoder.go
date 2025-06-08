package pgn

import (
	"math"
	"strings"

	"github.com/nedlir/chessencrypt/algorithm"
	"github.com/nedlir/chessencrypt/chess/board"
)

type PGNDecoder struct {
	algo algorithm.Algorithm
}

func NewPGNDecoder() PGNDecoder {
	return PGNDecoder{
		algo: algorithm.NewAlgorithm(),
	}
}

func (p *PGNDecoder) PGNToBytes(pgn string) []byte {
	pgnTokens := strings.Fields(pgn)
	firstBitValue := determineFirstBitValue(pgnTokens)
	moves := extractMovesFromPgn(pgnTokens)

	currentBlackMove := board.NewSquare(FIRST_BLACK_SQUARE)
	currentWhiteMove := board.NewSquare(FIRST_WHITE_SQUARE)

	var result []byte
	currentByte := byte(firstBitValue) << 7

	movesIndex := 0
	nextBlackMove := board.NewSquare(moves[movesIndex])
	nextWhiteMove := board.NewSquare(moves[movesIndex+1])

	currentByteIndex := 1

	for movesIndex+1 < len(moves) {
		isNewRow := isNewWhiteRow(currentWhiteMove, nextWhiteMove)

		if isNewRow {
			if currentByteIndex > 0 {
				result = append(result, currentByte)
			}
			currentByte = 0
			currentByteIndex = 0

			currentWhiteMove = nextWhiteMove
			currentBlackMove = nextBlackMove
			movesIndex += 2
			if movesIndex+1 < len(moves) {
				nextBlackMove = board.NewSquare(moves[movesIndex])
				nextWhiteMove = board.NewSquare(moves[movesIndex+1])
			}
			continue
		}

		var bit byte
		if currentByteIndex == nextWhiteMove.Column() {
			if isAssistanceMove(currentBlackMove, nextBlackMove) {
				bit = 0
			} else {
				bit = 1
			}
			currentBlackMove = nextBlackMove
			currentWhiteMove = nextWhiteMove
			movesIndex += 2
			if movesIndex+1 < len(moves) {
				nextBlackMove = board.NewSquare(moves[movesIndex])
				nextWhiteMove = board.NewSquare(moves[movesIndex+1])
			}
		} else {
			bit = 0
		}

		currentByte |= bit << (7 - currentByteIndex)

		if currentByteIndex == 7 {
			result = append(result, currentByte)
			currentByte = 0
			currentByteIndex = 0
		} else {
			currentByteIndex++
		}
	}

	if currentByteIndex > 0 {
		result = append(result, currentByte)
	}

	return result
}

func isNewWhiteRow(currentWhiteMove board.Square, nextWhiteMove board.Square) bool {
	return math.Abs(float64(currentWhiteMove.Row()-nextWhiteMove.Row())) >= 1
}

func isAssistanceMove(currentBlackMove board.Square, nextBlackMove board.Square) bool {
	return math.Abs(float64(currentBlackMove.Column()-nextBlackMove.Column())) > 1
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
	n := 0
	for i, token := range gameMoveTokens {
		if i%3 != 1 {
			gameMoveTokens[n] = token
			n++
		}
	}
	moves := gameMoveTokens[:n]

	if len(moves) > 0 {
		last := moves[len(moves)-1]
		if last == "1-0" || last == "0-1" || last == "1/2-1/2" || last == "*" {
			moves = moves[:len(moves)-1]
		}
	}

	for i, move := range moves {
		if len(move) > 0 && move[0] == 'Q' {
			moves[i] = move[1:]
		}
	}

	return moves
}
