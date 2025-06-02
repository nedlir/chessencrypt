package pgn

import (
	"fmt"
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

	fmt.Printf("tokens: %v\n firstBitValue: %d \n moves: %v", pgnTokens, firstBitValue, moves)

	return []byte{0b00000001}
}

func determineFirstBitValue(pgnTokens []string) int {
	firstFenValue := extractFirstFenValue(pgnTokens)

	fmt.Printf("\n\n\n\n\nfirstFenValue: %v\n\n\n\n", firstFenValue)

	if firstFenValue == board.WhiteKingOnA8 {
		return 0
	} else if firstFenValue == board.WhiteKingOnA7 {
		return 1
	}

	return -1

}

func extractFirstFenValue(pgnTokens []string) string {
	rawFen := pgnTokens[21]

	return rawFen[1:2] // skip first " of fen

}

func extractMovesFromPgn(pgnTokens []string) []string {
	gameMoveTokens := pgnTokens[28:]
	// Filter out move numbers (every 3rd token starting from index 1)
	// PGN format: "1... Qg8  2. Qb6  Qh8" -> we want ["Qg8", "Qb6", "Qh8"]
	n := 0
	for i, token := range gameMoveTokens {
		if i%3 != 1 {
			gameMoveTokens[n] = token
			n++
		}
	}

	moves := gameMoveTokens[:n]

	// Remove game result (last token "1/2-1/2")
	if len(moves) > 0 {
		last := moves[len(moves)-1]
		if last == "1-0" || last == "0-1" || last == "1/2-1/2" || last == "*" {
			moves = moves[:len(moves)-1]
		}
	}

	return moves
}
