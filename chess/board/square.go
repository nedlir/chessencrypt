package board

type Square struct {
	name   string
	row    int
	column int
}

func NewSquare(name string, row, column int) Square {
	return Square{
		name:   name,
		row:    row,
		column: column,
	}
}

func (s *Square) Name() string {
	return s.name
}

func (s *Square) SetName(name string) {
	s.name = name
}

func (s *Square) Row() int {
	return s.row
}

func (s *Square) Column() int {
	return s.column
}
