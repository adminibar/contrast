package parser

import (
	"github.com/dockpit/contrast/assert"
)

type AssertToFunc func(E) error

// Elements can be turned into assert
// functions that can test wether an other
// elements meet the exptation(s)
type E interface {
	Value() interface{}
	ToAssert(ats []*assert.Archetype) (AssertToFunc, error)
}

// Tables are compared against other
// tables
type T interface {
	All() map[string]E
	Set(string, E)
	Get(string) E
	Equals(T) error
}

// Parsers turn bytes into
// a linear Table of comparable elements
type P interface {
	Parse([]byte) (T, error)
}

// Element represents a value in a table
// that can converted to an assert
type Element struct {
	value interface{}
}

func NewElement(val interface{}) *Element {
	return &Element{val}
}

// Convert example value to string and ask the assert
// package to use it to generate a assertion function
func (example *Element) ToAssert(ats []*assert.Archetype) (AssertToFunc, error) {

	fn, err := assert.Parse(example.Value(), ats)
	if err != nil {
		return func(E) error { return err }, err
	}

	return func(actual E) error {
		return fn(example.Value(), actual.Value())
	}, nil
}

func (e *Element) Value() interface{} {
	return e.value
}
