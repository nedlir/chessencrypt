// // Pseudo code of the algorithm, need to check
// package algorithm

// const LEFT = "left"
// const RIGHT = "right"
// var whiteQueenDirection = LEFT

// var currentRow (serves as i)
// var currentColumn (serves as j)

// function setBits(matrix) {
// data []bytes{} ...
// b := NewGameMoves("black")
// w := NewGameMoves("w")

// isFirstBitToggled(get the first bit of first byte)
// 	fen = set the fen for this case
// else: fen =  the other fen

// for (matrix){

// nextBlackMove := generateNextMove(BLACK, b)
// b.add(nextBlackMove)

// nextWhiteMove := generateNextMove(WHITE, w)
// w.add (nextWhiteMove)
// }

// }

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

// func findNextBitToSet square () {
// 	this will loop through the matrix finding the next square needs to be set
// }

// func isReachableByBlackQueen() bool {
// }

// func isVerticalMove() bool {
// 	// this means that the Black Queen will just move down a row,
// 	// need to come up with a better variable name for this
// }

package algorithm

type Position struct {
	row    int
	column int
}

type Algorithm struct {
	bitMatrix [][]int
	position  Position
}

func NewAlgorithm(b [][]int) Algorithm {
	return Algorithm{bitMatrix: b, position: Position{row: 0, column: 0}}
}

func DetermineNextBlackMove() {}

func DetermineNextWhiteMove() {}
