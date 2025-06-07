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
	// Tokenize PGN and debug
	pgnTokens := strings.Fields(pgn)
	fmt.Printf("pgnTokens = %v\n", pgnTokens)

	// Determine initial bit from FEN
	firstBitValue := determineFirstBitValue(pgnTokens)
	fmt.Printf("firstBitValue = %d\n", firstBitValue)

	// Extract moves and debug
	moves := extractMovesFromPgn(pgnTokens)
	fmt.Printf("moves = %v\n\n", moves)

	// Initial positions
	currentBlackMove := board.NewSquare(FIRST_BLACK_SQUARE)
	currentWhiteMove := board.NewSquare(FIRST_WHITE_SQUARE)
	fmt.Printf("start Black at %s, White at %s\n\n",
		currentBlackMove.Name(), currentWhiteMove.Name(),
	)

	// Prepare result bitstring, seed first bit
	result := "0"
	if firstBitValue == 1 {
		result = "1"
	}
	fmt.Printf("initial result = %q\n\n", result)

	// Seed next moves
	movesIndex := 0
	nextBlackMove := board.NewSquare(moves[movesIndex])
	nextWhiteMove := board.NewSquare(moves[movesIndex+1])
	fmt.Printf("next Black = %s, next White = %s\n\n",
		nextBlackMove.Name(), nextWhiteMove.Name(),
	)

	currentByteIndex := 1

	for movesIndex+1 < len(moves) {
		fmt.Printf("\nLoop for moves %d/%d\n", movesIndex+2, len(moves))
		fmt.Printf("  Indexes -> movesIndex: %d, byteIndex: %d\n",
			movesIndex, currentByteIndex,
		)

		fmt.Printf("  Current Moves:\n")
		fmt.Printf("    White: %s (col %d, row %d)\n",
			currentWhiteMove.Name(), currentWhiteMove.Column(), currentWhiteMove.Row(),
		)
		fmt.Printf("    Black: %s (col %d)\n",
			currentBlackMove.Name(), currentBlackMove.Column(),
		)

		fmt.Printf("  Next Moves:\n")
		fmt.Printf("    White: %s (col %d, row %d)\n",
			nextWhiteMove.Name(), nextWhiteMove.Column(), nextWhiteMove.Row(),
		)
		fmt.Printf("    Black: %s (col %d)\n\n",
			nextBlackMove.Name(), nextBlackMove.Column(),
		)

		// Detect new White row
		isNewRow := isNewWhiteRow(currentWhiteMove, nextWhiteMove)
		fmt.Printf("  New White row detected: %v (from rank %d to %d)\n",
			isNewRow,
			currentWhiteMove.Row(),
			nextWhiteMove.Row(),
		)

		// Handle row break if needed
		if isNewRow {
			// Complete current byte with zeros if needed
			if currentByteIndex > 0 {
				for currentByteIndex < 8 {
					result += "0"
					currentByteIndex++
				}
			}
			result += " "
			currentByteIndex = 0

			// Update positions for new row
			currentWhiteMove = nextWhiteMove
			currentBlackMove = nextBlackMove
			movesIndex += 2
			if movesIndex+1 < len(moves) {
				nextBlackMove = board.NewSquare(moves[movesIndex])
				nextWhiteMove = board.NewSquare(moves[movesIndex+1])
			}
			fmt.Printf("  Row transition complete. New positions: White=%s, Black=%s\n",
				currentWhiteMove.Name(), currentBlackMove.Name())
			continue
		}

		// Determine bit for this move-pair
		var bit string
		if currentByteIndex == nextWhiteMove.Column() {
			if isAssistanceMove(currentBlackMove, nextBlackMove) {
				bit = "0"
			} else {
				bit = "1"
			}
			currentBlackMove = nextBlackMove
			currentWhiteMove = nextWhiteMove
			movesIndex += 2
			if movesIndex+1 < len(moves) {
				nextBlackMove = board.NewSquare(moves[movesIndex])
				nextWhiteMove = board.NewSquare(moves[movesIndex+1])
			}
		} else {
			bit = "0"
		}
		fmt.Printf("  appended bit = %q\n", bit)

		// Append bit and manage byte/delimiter
		result += bit
		if currentByteIndex == 7 {
			currentByteIndex = 0
		} else {
			currentByteIndex++
		}
		fmt.Printf("  updated result = %q\n", result)
	}

	// Complete final byte with zeros if needed
	if currentByteIndex > 0 && currentByteIndex < 8 {
		for currentByteIndex < 8 {
			result += "0"
			currentByteIndex++
		}
	}

	fmt.Printf("\nfinal bitstring = %q\n", result)
	return result
}

func (p *PGNDecoder) handleNewWhiteRow(
	currentWhiteMove, nextWhiteMove board.Square,
	result string, currentByteIndex int,
) (string, int) {
	fmt.Printf("Handling new row for white: %s -> %s\n",
		currentWhiteMove.Name(), nextWhiteMove.Name(),
	)
	result += "\n"
	currentByteIndex = 1
	return result, currentByteIndex
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
