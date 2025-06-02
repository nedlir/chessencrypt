package main

import (
	"fmt"

	"github.com/nedlir/chessencrypt/pgn"
)

func main() {

	matrix := []byte{
		0b11011010,
		// 0b10100101,
		// 0b10110110,
		// 0b00000001,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00101100,
		0b01101100,
	}

	fmt.Println("1")
	encoder := pgn.NewPGNEncoder()

	pgn1 := encoder.BytesToPgn(matrix, 1)

	fmt.Println("2")
	fmt.Printf("PGN:\n %v", pgn1)

	decoder := pgn.NewPGNDecoder()
	decoder.PGNToBytes(pgn1)

}
