package game

import (
	"fmt"

	bithandler "chessencryption/bitshandler"
	"chessencryption/chess/algorithm"
	. "chessencryption/chess/board"
)

func Run() {
	matrix := []byte{
		0b01011010,
		0b10100101,
		0b00110110,
		0b11001001,
		0b01101100,
		0b10010111,
	}

	bh := bithandler.NewBitHandler(matrix)
	mv := NewMoveValidator()
	algo := algorithm.NewAlgorithm(matrix, bh, mv)
	white := NewWhiteBoard()
	black := NewBlackBoard()

	fmt.Println("Bit matrix representation:")
	algo.PrintBitMatrix()
	fmt.Println()

	for {
		wMove, ok := algo.DetermineNextWhiteMove(&white)
		if !ok {
			fmt.Println("Game finished! All set bits have been processed.")
			return
		}
		white.AddMove(wMove)
		fmt.Printf(
			"White queen moved to %s (%d,%d) — value=%d\n",
			wMove.Name(), wMove.Row(), wMove.Column(), wMove.BinaryValue(),
		)

		// Black’s mirror move
		isAssist := wMove.BinaryValue() == 0
		bMove := algo.DetermineNextBlackMove(isAssist, &white)
		black.AddMove(bMove)
		fmt.Printf(
			"Black queen moved to %s (%d,%d)\n\n",
			bMove.Name(), bMove.Row(), bMove.Column(),
		)

		algo.SetCurrentSquare(&wMove)
	}
}
