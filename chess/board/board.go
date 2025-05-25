package board

const (
	WhiteQueenBoardRows int = 6
	WhiteQueenBoardCols int = 8

	BlackQueenBoardRows int = 1
	BlackQueenBoardCols int = 5
)

var WhiteQueenChessBoard = [WhiteQueenBoardRows][WhiteQueenBoardCols]string{
	{"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6"},
	{"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5"},
	{"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4"},
	{"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3"},
	{"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2"},
	{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1"},
}

var BlackQueenChessBoard = [BlackQueenBoardRows][BlackQueenBoardCols]string{
	{"d8", "e8", "f8", "g8", "h8"},
}
