package board

type Square struct {
	name   string
	row    int
	column int
}

func NewSquare(squareName string) Square {
	square, exists := BlackSquarePositions[squareName]
	if exists {
		return Square{
			name:   squareName,
			row:    square.Row,
			column: square.Col,
		}
	} else {
		square = WhiteSquarePositions[squareName]
	}
	return Square{
		name:   squareName,
		row:    square.Row,
		column: square.Col,
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
