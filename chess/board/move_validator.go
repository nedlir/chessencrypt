package board

import (
	"chessencryption/parsers/json"
	"fmt"
)

type queenValidDestinationsPerSquare map[Square]bool
type queenValidMovesMap map[Square]queenValidDestinationsPerSquare

type MovesValidator struct {
	possibleQueenMoves queenValidMovesMap
}

func NewMovesValidator() *MovesValidator {
	possibleQueenMoves, err := initQueenValidMoves("chess/data/queen_valid_moves.json")
	if err != nil {
		panic("failed to initialize valid moves: " + err.Error())
	}
	return &MovesValidator{possibleQueenMoves: possibleQueenMoves}
}

func (mv *MovesValidator) IsNextMoveValidMove(currentSquare, nextMove Square) bool {
	validDestinations, exists := mv.possibleQueenMoves[currentSquare]
	if !exists {
		return false
	}
	return validDestinations[nextMove]
}

func initQueenValidMoves(filepath string) (queenValidMovesMap, error) {
	data, err := json.LoadToMapFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load valid moves: %w", err)
	}

	vm := make(queenValidMovesMap)
	for key, valueMap := range data {
		sq := newSquareFromName(key)
		vm[sq] = make(queenValidDestinationsPerSquare)
		for destName := range valueMap {
			vm[sq][newSquareFromName(destName)] = true
		}
	}
	return vm, nil
}

func squareNameToRowCol(name string) (int, int) {
	if len(name) != 2 {
		return 0, 0
	}
	col := int(name[0] - 'a')
	row := 8 - int(name[1]-'0')
	return row, col
}

func newSquareFromName(name string) Square {
	row, col := squareNameToRowCol(name)
	return NewSquare(name, 0, row, col)
}
