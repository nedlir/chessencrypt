package main

import (
	"fmt"

	"github.com/nedlir/chessencrypt/pgn"
)

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
	matrix2 := []byte{
		0b01011010,
		0b10100101,
		0b10110110,
		0b00000001,
		0b00101100,
		0b01101100,
	}

	encoder := pgn.NewPGNEncoder()

	pgn1 := encoder.BytesToPgn(matrix, 1)

	fmt.Printf("PGN:\n %v", pgn1)

	fmt.Printf("PGN:\n %v", encoder.BytesToPgn(matrix2, 2))

}
