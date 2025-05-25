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

func (gm *GameMoves) IsNextMoveValidMove(nextMove string) bool {

	if gm.IsFirstMove() && gm.isBlack() && gm.isInvalidBlackFirstMove(nextMove) {
		return false
	}

	lastMove := gm.queenMoves[len(gm.queenMoves)-1]

	currentSquare := square(lastMove.square)
	nextSquare := square(nextMove)

	validDestinations, exists := gm.validMoves[currentSquare]
	if !exists {
		return false
	}

	isValid := validDestinations[nextSquare]

	return isValid
}

func (gm *GameMoves) IsFirstMove() bool {
	return len(gm.queenMoves) == 0
}

func (gm *GameMoves) isInvalidBlackFirstMove(firstMove string) bool {
	return firstMove != "Qg8" && firstMove != "Qh8"
}

func (gm *GameMoves) isBlack() bool {
	return gm.color == "black"
}

func (gm *GameMoves) isWhite() bool {
	return gm.color == "white"
}
