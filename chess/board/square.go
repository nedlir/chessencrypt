package board

type Square struct {
	name        string
	binaryValue int
}

func NewSquare(name string, binaryValue int) Square {
	return Square{
		name:        name,
		binaryValue: binaryValue,
	}
}

func NewSquareZero(name string) Square {
	return Square{
		name:        name,
		binaryValue: 0,
	}
}

func (s Square) Name() string {
	return s.name
}

func (s Square) BinaryValue() int {
	return s.binaryValue
}

func (s Square) Equals(other Square) bool {
	return s.name == other.name && s.binaryValue == other.binaryValue
}
