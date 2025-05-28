// // Pseudo code of the algorithm, need to check
// package algorithm

// func generateNextMove(color string, gm *GameMoves) {

// 	targetSquare := getLocationOfNextBitToSet()

// 	if color == BLACK {
// 		generateNextBlackMove(targetSquare)
// 	} else {
// 		generateNextWhiteMove(targetSquare)
// 	}
// }

// func generateNextWhiteMove(targetSquare square) square{
// 		isReachableByBlackQueen() {
// 			generateSingleRowMoveForWhite()
// 		}
// 		else{
// 			generateMultiRowMoveForWhite()
// 		}
// }

// func generateNextBlackMove(targetSquare square) square{
// 	isVerticalMove()
// 		{
// 			return generateMultiSquareMoveForBlack()
// 		}
// 		else{
// 			generateSingleSquareMoveForBlack()
// 		}
// 	}

// func getLocationOfBitToSet(will receive the bits matrix) square{
// 	check where is the location of next bit that needs to be drawn.
// 	will have to traverse the matrix to find the next bit that needs to be set.
// 	need to think if I want to traverse the matrix with a for loop
// 	maybe it is better to just keep the current column and index in the object.
// 	then to do a while loop until EOF.

// 	This funciton will just return matrix[i][j] with the name of the square
// }

// func isReachableByBlackQueen() bool {
// }

// func isVerticalMove() bool {
// 	// this means that the Black Queen will just move down a row,
// 	// need to come up with a better variable name for this
// }

package algorithm

import (
	. "chessencryption/chess/board"
	. "chessencryption/chess/fen"
	"fmt"
)

type Algorithm struct {
	bitMatrix     [][]int
	currentSquare Square
	nextSquare    Square
	moveValidator MoveValidator
}

func NewAlgorithm(b [][]int) Algorithm {
	return Algorithm{
		bitMatrix: b,
		currentSquare: NewSquare(
			"a6",
			0,
			NewPosition(0, 0),
		),
		nextSquare: NewSquare(
			"a6",
			0,
			NewPosition(0, 0),
		),
	}
}

func (a *Algorithm) PrintBitMatrix() {
	for row := 0; row < len(a.bitMatrix); row++ {
		for col := 0; col < len(a.bitMatrix[0]); col++ {
			fmt.Printf("%d ", a.bitMatrix[row][col])
		}
		fmt.Println()
	}
}

func (a *Algorithm) DetermineFEN(fen string) string {
	var firstBit int = a.bitMatrix[0][0]

	if firstBit == 0 {
		return FENZero
	} else {
		return FENOne
	}
}

func (a *Algorithm) FindNextBitToSet(cb *WhiteChessBoard) {
	currentPosition := a.currentSquare.Position()
	nextPosition := a.nextSquare.Position() // maybe add 3rd watcher

	for row := currentPosition.Row(); row < WhiteBoardRowsLength; row++ {
		for col := currentPosition.Column(); col < WhiteBoardColsLength; col++ {
			bit := a.bitMatrix[row][col]
			if bit == 1 {
				a.nextSquare = NewSquare(
					cb.Board()[row][col],
					1,
					NewPosition(row, col),
				)
				return
			}
		}
	}
}

func (a *Algorithm) CurrentSquare() Square {
	return a.currentSquare
}

func (a *Algorithm) NextSquare() Square {
	return a.nextSquare
}

func (a *Algorithm) DetermineNextBlackMove(isNextMoveAssistance bool) string {
	var nextSquare Square
	if isNextMoveAssistance {
		switch a.currentSquare.Name() {
		case "Qe8", "Qf8":
			nextSquare.SetName("Qh8")
		case "Qg8", "Qh8":
			nextSquare.SetName("Qe8")
		}
	} else { // next move is supposed to mark the bit
		switch a.currentSquare.Name() {
		case "Qe8":
			nextSquare.SetName("Qf8")
		case "Qf8":
			nextSquare.SetName("Qg8")
		case "Qg8":
			nextSquare.SetName("Qh8")
		case "Qh8":
			nextSquare.SetName("Qg8")
		}

	}
	return "f"
}

func (a *Algorithm) DetermineNextWhiteMove() string {
	// isValidNextWhiteMove()
	return "a"
}
