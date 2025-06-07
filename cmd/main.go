package main

import (
	"fmt"

	"github.com/nedlir/chessencrypt/chess/pgn"
	"github.com/nedlir/chessencrypt/utils/fileshandler"
)

func main() {
	encoder := pgn.NewPGNEncoder()

	fileContent, err := fileshandler.ReadFile("in.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	pgn1 := encoder.BytesToPgn(fileContent, 1)

	fmt.Printf("file content: %08b\n", fileContent)

	fileshandler.WriteFile("out.pgn", []byte(pgn1))

	decoder := pgn.NewPGNDecoder()
	decodedBits := decoder.PGNToBytes(pgn1)

	// Convert string bits to bytes
	result := make([]byte, 6)
	currentByte := byte(0)
	currentRow := 0

	for i := 0; i < len(decodedBits); i++ {
		if decodedBits[i] == '1' {
			currentByte |= 1 << (7 - (i % 8))
		}
		if (i+1)%8 == 0 {
			result[currentRow] = currentByte
			currentRow++
			currentByte = 0
		}
	}

	fmt.Printf("\n\nDecoded bytes in binary:\n")
	for i, b := range result {
		fmt.Printf("Row %d: %08b\n", i, b)
	}

	fmt.Println("\nfinished")
}
