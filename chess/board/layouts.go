package board

const (
	WhiteBoardRowsLength = 6
	WhiteBoardColsLength = 8
	BlackBoardColsLength = 5
)

type WhiteBoardMatrix [WhiteBoardRowsLength][WhiteBoardColsLength]string
type BlackBoardMatrix [BlackBoardColsLength]string

type Position struct {
	Row int
	Col int
}

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

var WhiteSquarePositions = map[string]Position{
	"a6": {0, 0}, "b6": {0, 1}, "c6": {0, 2}, "d6": {0, 3}, "e6": {0, 4}, "f6": {0, 5}, "g6": {0, 6}, "h6": {0, 7},
	"a5": {1, 0}, "b5": {1, 1}, "c5": {1, 2}, "d5": {1, 3}, "e5": {1, 4}, "f5": {1, 5}, "g5": {1, 6}, "h5": {1, 7},
	"a4": {2, 0}, "b4": {2, 1}, "c4": {2, 2}, "d4": {2, 3}, "e4": {2, 4}, "f4": {2, 5}, "g4": {2, 6}, "h4": {2, 7},
	"a3": {3, 0}, "b3": {3, 1}, "c3": {3, 2}, "d3": {3, 3}, "e3": {3, 4}, "f3": {3, 5}, "g3": {3, 6}, "h3": {3, 7},
	"a2": {4, 0}, "b2": {4, 1}, "c2": {4, 2}, "d2": {4, 3}, "e2": {4, 4}, "f2": {4, 5}, "g2": {4, 6}, "h2": {4, 7},
	"a1": {5, 0}, "b1": {5, 1}, "c1": {5, 2}, "d1": {5, 3}, "e1": {5, 4}, "f1": {5, 5}, "g1": {5, 6}, "h1": {5, 7},
}

var BlackSquarePositions = map[string]Position{
	"d8": {0, 0},
	"e8": {0, 1},
	"f8": {0, 2},
	"g8": {0, 3},
	"h8": {0, 4},
}
