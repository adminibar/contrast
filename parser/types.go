package parser

type AssertToFunc func(E) error

// Elements can be turned into assert
// functions that can test wether an other
// elements meet the exptation(s)
type E interface {
	Value() interface{}
	ToAssert() (AssertToFunc, error)
}

// Tables are compared against other
// tables
type T interface {
	All() map[string]E
	Set(string, E)
	Get(string) E

	//expect this table to have at least the
	//elements of the other table but its ok
	//if it has any extra
	AtLeast(T) error
}

// Parsers turn bytes into
// a linear Table of comparable elements
type P interface {
	Parse([]byte) (T, error)
}
