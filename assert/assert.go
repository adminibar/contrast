package assert

import (
	"github.com/stretchr/testify/assert"
)

type AssertFunc func(expected, actual interface{}) error

//parse an example value to determine how the actual value
//will be asserted, (e.g. strict equal, type only)
func Parse(example interface{}) (AssertFunc, error) {

	//@todo assert example to type and analyse content
	//to determine comparision strategy (e.g strict equal)

	return StrictEqualStrategy(), nil
}

//Value should be exactly equal to the example
func StrictEqualStrategy() AssertFunc {
	return func(example, actual interface{}) error {
		t := newTest()
		assert.Equal(t, example, actual)
		return t.Error()
	}
}
