package board

import . "chessencryption/chess/constants"

type Position struct {
	row    Row
	column Column
	square Square
}

func NewPosition(row Row, column Column, square Square) Position {
	return Position{
		row:    row,
		column: column,
		square: square,
	}
}

func (p Position) Row() Row {
	return p.row
}

func (p Position) Column() Column {
	return p.column
}

func (p Position) Square() Square {
	return p.square
}
