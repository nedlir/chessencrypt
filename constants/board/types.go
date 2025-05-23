package board

type Square string

type SquareValue int

const (
	Zero SquareValue = 0
	One  SquareValue = 1
)

type Board [][]Square

type Matrix [][]SquareValue
