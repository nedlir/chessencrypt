package pgn

import (
	"fmt"
	"strings"

	"chessencryption/algorithm"
	"chessencryption/bitshandler"
	"chessencryption/chess/board"
)

const WHITE_QUEEN_MOVE_ZERO = ".."
const FIRST_WHITE_SQUARE = "a6"
const FIRST_BLACK_SQUARE = "f8"
const END_OF_GAME_DRAW = "1/2-1/2"

func Run(matrix []byte) {

	bitsHandler := bitshandler.NewBitHandler()
	algo := algorithm.NewAlgorithm(matrix)

	var pgnBoard strings.Builder
	fluff := `[Event "?"]
[Site "?"]
[Date "????.??.??"]
[Round "?"]
[White "?"]
[Black "?"]
[Result "1/2-1/2"]
[WhiteELO "?"]
[BlackELO "?"]
[SetUp "1"]
`
	pgnBoard.WriteString(fluff)

	bitsHandler.UpdateMatrix(matrix)
	fmt.Printf("Bit matrix representation: \n %b \n", matrix)
	squaresToMark := bitsHandler.AllSetBits()
	fmt.Printf("Found bits positions: %v", squaresToMark)

	if bitsHandler.IsFirstBitZero() {
		pgnBoard.WriteString(fmt.Sprintf(`[FEN "%v"]`+"\n", board.FENZero))
	} else {
		pgnBoard.WriteString(fmt.Sprintf(`[FEN "%v"]`+"\n", board.FENOne))
	}

	pgnBoard.WriteString(fmt.Sprintf("1.%s ", WHITE_QUEEN_MOVE_ZERO)) // "1..."

	currentWhiteSquare := board.NewSquare(FIRST_WHITE_SQUARE, 0, 0, 0)
	nextWhiteSquare := currentWhiteSquare
	currentBlackSquare := board.NewSquare(FIRST_BLACK_SQUARE, 0, 0, 5)

	isAssist := false

	squareIndex := 0
	pgnMoveIndex := 2

	for squareIndex < len(squaresToMark) {
		nextWhiteSquare, isAssist = algo.DetermineNextWhiteMove(currentWhiteSquare, squaresToMark[squareIndex])
		currentWhiteSquare = nextWhiteSquare

		if !isAssist {
			squareIndex++
		}

		currentBlackSquare = algo.DetermineNextBlackMove(isAssist, currentBlackSquare)

		//add black move
		pgnBoard.WriteString(fmt.Sprintf("Q%v  ", currentBlackSquare.Name()))
		pgnBoard.WriteString(fmt.Sprintf("%d. ", pgnMoveIndex))
		pgnMoveIndex++

		// add white move
		pgnBoard.WriteString(fmt.Sprintf("Q%v  ", nextWhiteSquare.Name()))

	}

	pgnBoard.WriteString(END_OF_GAME_DRAW)

	fmt.Printf("PGN:\n%v", pgnBoard.String())
}
