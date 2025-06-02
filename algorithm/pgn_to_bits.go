package algorithm

// func calculatePositionalDistance(currentSquare string, nextSquare string) int {

// 	return int(nextSquare - currentSquare)
// }

// func DetermineNextSetBit() {
// 	nextBitToSet
// }

func calculatePositionalDistance(currentSquare, nextSquare string) int {
	return int(nextSquare[0]) - int(currentSquare[0])
}
