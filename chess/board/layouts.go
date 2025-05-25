package board

import . "chessencryption/chess/constants"

const (
	WhiteBoardRows = 6
	WhiteBoardCols = 8
	BlackBoardCols = 5
)

var (
	BLACK_START_POSITION = NewPosition(0, 2, "f8")
	WHITE_START_POSITION = NewPosition(0, 0, "a6")
)

type WhiteBoardMatrix [WhiteBoardRows][WhiteBoardCols]Square
type BlackBoardMatrix [BlackBoardCols]Square

type BoardMatrix interface {
	WhiteBoardMatrix | BlackBoardMatrix
}

var WhiteQueenLayout = WhiteBoardMatrix{
	{"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6"},
	{"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5"},
	{"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4"},
	{"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3"},
	{"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2"},
	{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1"},
}

var BlackQueenLayout = BlackBoardMatrix{
	"d8", "e8", "f8", "g8", "h8",
}
