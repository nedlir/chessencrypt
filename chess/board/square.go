package board

type Square struct {
	name        string
	binaryValue int
	row         int
	column      int
}

func NewSquare(name string, binaryValue, row, column int) Square {
	return Square{
		name:        name,
		binaryValue: binaryValue,
		row:         row,
		column:      column,
	}
}

func NewSquareZero(s Square) Square {
	return Square{
		name:        s.name,
		binaryValue: 0,
		row:         s.row,
		column:      s.column,
	}
}

func (s *Square) Name() string {
	return s.name
}

func (s *Square) SetName(name string) {
	s.name = name
}

func (s *Square) BinaryValue() int {
	return s.binaryValue
}

func (s *Square) SetBinaryValue(val int) {
	s.binaryValue = val
}

func (s *Square) Row() int {
	return s.row
}

func (s *Square) Column() int {
	return s.column
}
