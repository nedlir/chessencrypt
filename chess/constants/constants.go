package constants

type (
	Color  string
	Square string
	Row    int
	Column int
)

const (
	BLACK                       Color  = "black"
	WHITE                       Color  = "black"
	BLACK_QUEEN_STARTING_SQUARE Square = "Qf8"
	WHITE_QUEEN_STARTING_SQUARE Square = "Qa6"
	// TODO: need to think later how to represent the start of the game for black.
	WHITE_QUEEN_MOVE_ZERO Square = "..."
)
