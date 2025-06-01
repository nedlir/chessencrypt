package main

import "github.com/nedlir/chessencrypt/chess/pgn"

// import "chessencryption/chess/game"

func main() {

	matrix := []byte{
		0b01011010,
		// 0b10100101,
		// 0b10110110,
		// 0b00000001,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00101100,
		0b01101100,
	}

	pgn.Run(matrix)

}
