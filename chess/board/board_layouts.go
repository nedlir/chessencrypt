package board

const (
	WhiteBoardRows = 6
	WhiteBoardCols = 8
	BlackBoardCols = 5
)

type WhiteBoardMatrix [WhiteBoardRows][WhiteBoardCols]Square
type BlackBoardMatrix [BlackBoardCols]Square

var WhiteQueenLayout = WhiteBoardMatrix{
	{NewSquareZero("a6"), NewSquareZero("b6"), NewSquareZero("c6"), NewSquareZero("d6"), NewSquareZero("e6"), NewSquareZero("f6"), NewSquareZero("g6"), NewSquareZero("h6")},
	{NewSquareZero("a5"), NewSquareZero("b5"), NewSquareZero("c5"), NewSquareZero("d5"), NewSquareZero("e5"), NewSquareZero("f5"), NewSquareZero("g5"), NewSquareZero("h5")},
	{NewSquareZero("a4"), NewSquareZero("b4"), NewSquareZero("c4"), NewSquareZero("d4"), NewSquareZero("e4"), NewSquareZero("f4"), NewSquareZero("g4"), NewSquareZero("h4")},
	{NewSquareZero("a3"), NewSquareZero("b3"), NewSquareZero("c3"), NewSquareZero("d3"), NewSquareZero("e3"), NewSquareZero("f3"), NewSquareZero("g3"), NewSquareZero("h3")},
	{NewSquareZero("a2"), NewSquareZero("b2"), NewSquareZero("c2"), NewSquareZero("d2"), NewSquareZero("e2"), NewSquareZero("f2"), NewSquareZero("g2"), NewSquareZero("h2")},
	{NewSquareZero("a1"), NewSquareZero("b1"), NewSquareZero("c1"), NewSquareZero("d1"), NewSquareZero("e1"), NewSquareZero("f1"), NewSquareZero("g1"), NewSquareZero("h1")},
}

var BlackQueenLayout = BlackBoardMatrix{
	NewSquareZero("d8"), NewSquareZero("e8"), NewSquareZero("f8"), NewSquareZero("g8"), NewSquareZero("h8"),
}
