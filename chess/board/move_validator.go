package board

import (
	"chessencryption/parsers/json"
	"fmt"
)

type queenValidMovesMap map[string][]string

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

func (mv *MovesValidator) IsNextMoveValidMove(currentSquare Square, nextMove Square) bool {
	fmt.Println("\nIsNextMoveValidMove:")
	fmt.Printf("current Square: %v\n next Square: %v \n", currentSquare, nextMove)

	currentSquareName := currentSquare.Name()
	nextMoveName := nextMove.Name()

	validDestinations, exists := mv.possibleQueenMoves[currentSquareName]
	if !exists {
		fmt.Printf("No moves found for square: %s\n", currentSquareName)
		return false
	}

	for _, validMove := range validDestinations {
		if validMove == nextMoveName {
			return true
		}
	}

	return false
}

func initQueenValidMoves(filepath string) (queenValidMovesMap, error) {
	data, err := json.LoadToMapFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load valid moves: %w", err)
	}

	vm := make(queenValidMovesMap)
	for squareName, valueMap := range data {
		var destinations []string

		for destName := range valueMap {
			destinations = append(destinations, destName)
		}

		vm[squareName] = destinations
	}

	return vm, nil
}
