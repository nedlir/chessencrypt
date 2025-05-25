package moves

import (
	"chessencryption/parsers"
	"fmt"
	"strings"
)

type square string
type validMovesPerSquare map[square](bool)
type validMoves map[square](validMovesPerSquare)

type GameMoves struct {
	color      string
	queenMoves []queenMove
	validMoves validMoves
}

func NewGameMoves(color string) *GameMoves {
	color = strings.ToLower(color)
	if color != "white" && color != "black" {
		panic("invalid color: must be 'white' or 'black'")
	}

	validMoves, err := initValidMoves("game/data/queen_valid_moves.json")
	if err != nil {
		panic("failed to initialize valid moves: " + err.Error())

	}

	gm := &GameMoves{
		color:      color,
		queenMoves: []queenMove{},
		validMoves: validMoves,
	}

	return gm
}

// func AddMove(move string) {

// }

// func isValidMove(move string) bool {

// }

func initValidMoves(filepath string) (validMoves, error) {
	parser := parsers.NewJSONParser()

	data, err := parser.LoadToMapFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load valid moves: %w", err)
	}

	vm := make(validMoves)
	for key, valueMap := range data {
		sq := square(key)
		vm[sq] = make(validMovesPerSquare)
		for value := range valueMap {
			vm[sq][square(value)] = true
		}
	}

	return vm, nil
}
