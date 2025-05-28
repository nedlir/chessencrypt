package board

const (
	WhiteBoardRowsLength = 6
	WhiteBoardColsLength = 8
	BlackBoardColsLength = 5
)

type WhiteBoardMatrix [WhiteBoardRowsLength][WhiteBoardColsLength]string
type BlackBoardMatrix [BlackBoardColsLength]string

var WhiteQueenLayout WhiteBoardMatrix = [WhiteBoardRowsLength][WhiteBoardColsLength]string{
	{"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6"},
	{"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5"},
	{"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4"},
	{"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3"},
	{"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2"},
	{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1"},
}

var BlackQueenLayout BlackBoardMatrix = [BlackBoardColsLength]string{
	"d8", "e8", "f8", "g8", "h8",
}
