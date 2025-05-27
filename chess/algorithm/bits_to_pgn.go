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
	. "chessencryption/chess/fen"
	"fmt"
)

type Position struct {
	row    int
	column int
}

func (p *Position) Row() int {
	return p.row
}

func (p *Position) Column() int {
	return p.column
}

type Algorithm struct {
	bitMatrix       [][]int
	currentPosition Position
	nextPosition    Position
	scanPosition    Position
}

func NewAlgorithm(b [][]int) Algorithm {
	return Algorithm{bitMatrix: b, currentPosition: Position{0, 0}, nextPosition: Position{0, 0}, scanPosition: Position{0, 0}}
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

// For now will use brute-force -
// queen will finish 1 line,
// then queen will go down a line
// then queen will move to square [row][col=0] and will continue scanning from there.
// meaning I will add 2 assistance moves for the queen and total of 4 in general
// TODO: optimize this by finding shortest path....

func (a *Algorithm) FindNextBitToSet() {
	rows, cols := len(a.bitMatrix), len(a.bitMatrix[0])
	for r := a.scanPosition.row; r < rows; r++ {
		startC := a.scanPosition.column
		if r > a.scanPosition.row {
			startC = 0
		}
		for c := startC; c < cols; c++ {
			if a.bitMatrix[r][c] == 1 {
				// record it
				a.nextPosition = Position{r, c}
				// advance the scan cursor for next time
				if c+1 < cols {
					a.scanPosition = Position{r, c + 1}
				} else {
					a.scanPosition = Position{r + 1, 0}
				}
				return
			}
		}
	}
	// no more bits
	a.scanPosition = Position{rows, 0}
}

func (a *Algorithm) NextPosition() Position {
	return a.nextPosition
}

func (a *Algorithm) CurrentPosition() Position {
	return a.currentPosition
}

func (a *Algorithm) DetermineNextBlackMove() string {
	// isValidNextBlackMove()
	return "f"
}

func (a *Algorithm) DetermineNextWhiteMove() string {
	// isValidNextWhiteMove()
	return "a"
}

func (a *Algorithm) IsGameFinished() bool {
	return a.scanPosition.row >= len(a.bitMatrix)
}

// func (cb *WhiteChessBoard) IsNextMoveValidMove(nextMove Square) bool {
// 	return cb.moveValidator.IsNextMoveValidMove(cb.queenMoves, nextMove)
// }

func (a *Algorithm) ApplyNextMove() {
	a.currentPosition = a.nextPosition
}
