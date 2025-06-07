package pgn

import (
	"fmt"
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

func (p *PGNDecoder) PGNToBytes(pgn string) string {
	pgnTokens := strings.Fields(pgn)
	fmt.Printf("pgnTokens = %v\n", pgnTokens)

	firstBitValue := determineFirstBitValue(pgnTokens)
	fmt.Printf("firstBitValue = %d\n", firstBitValue)

	moves := extractMovesFromPgn(pgnTokens)
	fmt.Printf("moves = %v\n\n", moves)

	// Split moves into bytes (each byte needs 8 bits)
	var bytesMoves [][]string
	for i := 0; i < len(moves); i += 8 {
		end := i + 8
		if end > len(moves) {
			end = len(moves)
		}
		bytesMoves = append(bytesMoves, moves[i:end])
	}

	// Process each byte's moves
	result := ""
	for byteIndex, byteMoves := range bytesMoves {
		// Initial positions for this byte
		currentBlackMove := board.NewSquare(FIRST_BLACK_SQUARE)
		currentWhiteMove := board.NewSquare(FIRST_WHITE_SQUARE)
		fmt.Printf("Processing byte %d, starting at %s, %s\n",
			byteIndex, currentBlackMove.Name(), currentWhiteMove.Name(),
		)

		// Set first bit
		if byteIndex == 0 {
			if firstBitValue == 1 {
				result += "1"
			} else {
				result += "0"
			}
		} else {
			result += "0" // Subsequent bytes start with 0
		}

		// Process moves for this byte
		currentByteIndex := 1
		movesIndex := 0

		if len(byteMoves) > 1 {
			nextBlackMove := board.NewSquare(byteMoves[movesIndex])
			nextWhiteMove := board.NewSquare(byteMoves[movesIndex+1])

			for movesIndex+1 < len(byteMoves) && currentByteIndex < 8 {
				fmt.Printf("byte %d: movesIndex=%d | byteIndex=%d\n",
					byteIndex, movesIndex, currentByteIndex,
				)
				fmt.Printf("current White = %s(col %d), Black = %s(col %d)\n",
					currentWhiteMove.Name(), currentWhiteMove.Column(),
					currentBlackMove.Name(), currentBlackMove.Column(),
				)
				fmt.Printf("next White = %s(col %d), Black = %s(col %d)\n",
					nextWhiteMove.Name(), nextWhiteMove.Column(),
					nextBlackMove.Name(), nextBlackMove.Column(),
				)

				var bit string
				if currentByteIndex == nextWhiteMove.Column() {
					if isAssistanceMove(currentBlackMove, nextBlackMove) {
						bit = "0assistance"
					} else {
						bit = "1"
					}
					currentBlackMove = nextBlackMove
					currentWhiteMove = nextWhiteMove

					movesIndex += 2
					if movesIndex+1 < len(byteMoves) {
						nextBlackMove = board.NewSquare(byteMoves[movesIndex])
						nextWhiteMove = board.NewSquare(byteMoves[movesIndex+1])
					}
				} else {
					bit = "0"
				}
				fmt.Printf("appended bit = %q\n", bit)

				result += bit
				currentByteIndex++
			}
		}

		// Fill remaining bits with 0s if needed
		for currentByteIndex < 8 {
			result += "0"
			currentByteIndex++
		}
	}

	fmt.Printf("final bitstring = %q\n", result)
	return result
}

func isNewWhiteRow(currentWhiteMove board.Square, nextWhiteMove board.Square) bool {
	return math.Abs(float64(currentWhiteMove.Row()-nextWhiteMove.Row())) >= 1
}

func isAssistanceMove(currentBlackMove board.Square, nextBlackMove board.Square) bool {
	fmt.Println("Abs(float64(currentBlackMove.Column()-nextBlackMove.Column())):")
	res := math.Abs(float64(currentBlackMove.Column() - nextBlackMove.Column()))
	fmt.Println(res)
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
	return rawFen[1:2] // skip leading quote
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
