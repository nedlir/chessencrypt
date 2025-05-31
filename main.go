package main

import (
	"chessencryption/chess/game"
	"math/bits"
)

// import "chessencryption/chess/game"

func main() {

	game.Run()

	// myByte := byte(0b00110110)
	// result := firstSetBitFromLeft(myByte)

	// fmt.Printf("Byte is: %08b, first appearance of One from left at index: %d", myByte, result)
}

func firstSetBitFromLeft(b byte) int {
	if b == 0 {
		return -1 // no 1 bits
	}

	// bits.LeadingZeros8 counts zeros from the left
	leadingZeros := bits.LeadingZeros8(b)

	// The first 1 bit from left is at position equal to leading zeros count
	return leadingZeros
}
