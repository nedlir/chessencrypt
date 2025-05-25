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

func (p Position) GetCoordinates() (Row, Column) {
	return p.row, p.column
}

func (p Position) String() string {
	return string(p.square)
}

func (p Position) IsValid() bool {
	return p.square != ""
}

func (p Position) Equals(other Position) bool {
	return p.row == other.row && p.column == other.column && p.square == other.square
}

func (p *Position) UpdatePosition(row Row, column Column, square Square) {
	p.row = row
	p.column = column
	p.square = square
}
