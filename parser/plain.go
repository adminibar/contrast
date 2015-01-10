package parser

import (
	"strings"
	"unicode"

	"github.com/dockpit/contrast/assert"
)

// A plain text element
type PlainE struct {
	value interface{}
}

func NewPlainE(val interface{}) *PlainE {
	return &PlainE{val}
}

// Convert example value to string and ask the assert
// package to use it to generate a assertion function
func (example *PlainE) ToAssert() (AssertToFunc, error) {

	fn, err := assert.Parse(example)
	if err != nil {
		return func(E) error { return err }, err
	}

	return func(actual E) error {
		return fn(example.Value(), actual.Value())
	}, nil
}

func (e *PlainE) Value() interface{} {
	return e.value
}

// A table that should only contain the
// single plain text value
type PlainT struct {
	values map[string]E
}

func newPlainT() *PlainT {
	return &PlainT{
		values: map[string]E{},
	}
}

func (t *PlainT) All() map[string]E {
	return t.values
}

func (t *PlainT) Set(key string, e E) {
	t.values[key] = e
}

func (t *PlainT) Get(key string) E {
	return t.values[key]
}

func (t *PlainT) AtLeast(ex T) error {

	return nil
}

// For parsing byte arrays that just hold
// plain text but will ignore trailing whitespace
type Plain struct{}

func NewPlain() *Plain {
	return &Plain{}
}

func (p *Plain) Parse(data []byte) (T, error) {
	t := newPlainT()

	//convert to string and remove spaces according to unicode
	str := strings.TrimRightFunc(string(data), unicode.IsSpace)

	//creat element
	e := NewPlainE(str)

	//set it as single table value
	t.Set(".0", e)

	return t, nil
}
