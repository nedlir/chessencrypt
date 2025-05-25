package main

import (
	"chessencryption/game/moves"
	"fmt"
)

func main() {
	game := moves.NewGameMoves("black")

	fmt.Print(game)

}
