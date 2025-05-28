package board

type Position struct {
	row    int
	column int
}

func (p *Position) Row() int {
	return p.row
}

func (p *Position) Column() int {
	return p.column
}
func NewPosition(row, column int) Position {
	return Position{
		row:    row,
		column: column,
	}
}
