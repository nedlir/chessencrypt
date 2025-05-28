package board

type Square struct {
	name        string
	binaryValue int
	position    Position
}

func NewSquare(name string, binaryValue int, position Position) Square {
	return Square{
		name:        name,
		binaryValue: binaryValue,
		position:    position,
	}
}

func NewSquareZero(square Square) Square {
	return Square{
		name:        square.name,
		binaryValue: 0,
		position:    Position{square.position.row, square.position.column},
	}
}

func (s *Square) Name() string {
	return s.name
}

func (s *Square) Position() Position {
	return s.position
}

func (s *Square) SetName(name string) {
	s.name = name
}

func (s *Square) BinaryValue() int {
	return s.binaryValue
}
