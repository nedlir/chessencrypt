package validators

import (
	. "chessencryption/chess/constants"
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

func (mv *MoveValidator) IsNextMoveValidMove(queenMoves []Square, color Color, nextMove Square) bool {
	if len(queenMoves) == 0 && color == BLACK && (nextMove != Square("Qg8") && nextMove != Square("Qh8")) {
		return false
	}

	if len(queenMoves) == 0 {
		return false
	}

	lastMove := queenMoves[len(queenMoves)-1]

	validDestinations, exists := mv.possibleQueenMoves[lastMove]
	if !exists {
		return false
	}

	return validDestinations[nextMove]
}

func initQueenValidMoves(filepath string) (queenValidMovesMap, error) {
	parser := parsers.NewJSONParser()

	data, err := parser.LoadToMapFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load valid moves: %w", err)
	}

	vm := make(queenValidMovesMap)
	for key, valueMap := range data {
		sq := Square(key)
		vm[sq] = make(queenValidDestinationsPerSquare)
		for value := range valueMap {
			vm[sq][Square(value)] = true
		}
	}

	return vm, nil
}
