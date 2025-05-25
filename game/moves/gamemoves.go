package moves

import (
	"chessencryption/parsers"
	"fmt"
	"strings"
)

const (
	BLACK_QUEEN_STARTING_SQUARE square = "Qf8"
	WHITE_QUEEN_STARTING_SQUARE square = "Qa6"
	BLACK                              = "black"
	WHITE                              = "white"
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
	if color != WHITE && color != BLACK {
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

func (gm *GameMoves) IsNextMoveValidMove(nextMove square) bool {

	if gm.isFirstMove() && gm.isBlack() && gm.isInvalidBlackFirstMove(nextMove) {
		return false
	}

	lastMove := gm.queenMoves[len(gm.queenMoves)-1]

	lastMoveSquare := lastMove.square
	nextMoveSquare := nextMove

	validDestinations, exists := gm.validMoves[lastMoveSquare]
	if !exists {
		return false
	}

	isValid := validDestinations[nextMoveSquare]

	return isValid
}

func (gm *GameMoves) AddMove(move square) {
	if !gm.IsNextMoveValidMove(move) {
		panic("Move inserted is invalid")
	}
	gm.queenMoves = append(gm.queenMoves, newQueenMove(move))
}

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

func (gm *GameMoves) isFirstMove() bool {
	return len(gm.queenMoves) == 0
}

func (gm *GameMoves) isInvalidBlackFirstMove(firstMove square) bool {
	return firstMove != "Qg8" && firstMove != "Qh8"
}

func (gm *GameMoves) isBlack() bool {
	return gm.color == BLACK
}

func (gm *GameMoves) isWhite() bool {
	return gm.color == WHITE
}
