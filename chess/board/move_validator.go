package board

import (
	"chessencryption/parsers"
	"fmt"
)

type queenValidDestinationsPerSquare map[Square]bool
type queenValidMovesMap map[Square]queenValidDestinationsPerSquare

type MoveValidator struct {
	possibleQueenMoves queenValidMovesMap
}

func NewMoveValidator() *MoveValidator {
	possibleQueenMoves, err := initQueenValidMoves("chess/data/queen_valid_moves.json")
	if err != nil {
		panic("failed to initialize valid moves: " + err.Error())
	}

	return &MoveValidator{
		possibleQueenMoves: possibleQueenMoves,
	}
}

func (mv *MoveValidator) IsNextMoveValidMove(currentSquare, nextMove Square) bool {
	validDestinations, exists := mv.possibleQueenMoves[currentSquare]
	if !exists {
		return false
	}

	return validDestinations[nextMove]
}

func squareNameToPosition(squareName string) Position {
	if len(squareName) != 2 {
		return NewPosition(0, 0)
	}

	col := int(squareName[0] - 'a')

	row := 8 - int(squareName[1]-'0')

	return NewPosition(row, col)
}

func newSquareFromName(squareName string) Square {
	position := squareNameToPosition(squareName)
	return NewSquare(squareName, 0, position)
}

func initQueenValidMoves(filepath string) (queenValidMovesMap, error) {
	parser := parsers.NewJSONParser()

	data, err := parser.LoadToMapFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load valid moves: %w", err)
	}

	vm := make(queenValidMovesMap)
	for key, valueMap := range data {
		sq := newSquareFromName(key)
		vm[sq] = make(queenValidDestinationsPerSquare)

		for value := range valueMap {
			vm[sq][newSquareFromName(value)] = true
		}
	}

	return vm, nil
}
