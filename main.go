package main

import (
	"chessencryption/chess/moves"
)

func main() {
	var blackQueenMoves []moves.QueenMove
	first_black_move := moves.NewQueenMove("-")
	blackQueenMoves = append(blackQueenMoves, first_black_move)
}
