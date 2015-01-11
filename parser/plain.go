package parser

import (
	"strings"
	"unicode"

	"github.com/dockpit/contrast/assert"
)

// A table that should only contain the
// single plain text value
type PlainT struct {
	values     map[string]E
	archetypes []*assert.Archetype
}

func newPlainT(ats []*assert.Archetype) *PlainT {
	if ats == nil {
		ats = []*assert.Archetype{}
	}

	return &PlainT{
		values:     map[string]E{},
		archetypes: ats,
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

	//the tables only value into an assert
	actual := t.Get(".0")
	fn, err := actual.ToAssert(t.archetypes)
	if err != nil {
		return err
	}

	return fn(ex.Get(".0"))
}

// For parsing byte arrays that just hold
// plain text but will ignore trailing whitespace
// as defined by unicode
type Plain struct{}

func NewPlain() *Plain {
	return &Plain{}
}

//if ats is nill, an empty list of archetypes is used
func (p *Plain) Parse(data []byte, ats []*assert.Archetype) (T, error) {
	t := newPlainT(ats)

	//convert to string and remove spaces according to unicode
	str := strings.TrimRightFunc(string(data), unicode.IsSpace)

	//creat element
	e := NewElement(str)

	//set it as single table value
	t.Set(".0", e)

	return t, nil
}