package strategy

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

type Strategy func() AssertFunc
type AssertFunc func(expected, actual interface{}) error

//Mapped list of available strategies
var AvailableStrategies = map[string]Strategy{
	"strict": StrictEqualStrategy,
	"type":   TypeEqualStrategy,
}

//parse an example value to determine how the actual value
//will be asserted, (e.g. strict equal, type only)
func Parse(example interface{}, ats []*Archetype) (AssertFunc, error) {
	var err error
	//ask each arhectype if the example value fits
	//if it deos the archetype defines the assertion
	//strategy
	strat := StrictEqualStrategy
	for _, at := range ats {

		if at.Fits(example) {
			strat, err = at.ToStrategy()
			if err != nil {
				return nil, fmt.Errorf("Error formulating strategy: %s", err)
			}

			break
		}
	}

	return strat(), nil
}

//Actual value should be exactly equal to the example
func StrictEqualStrategy() AssertFunc {
	return func(example, actual interface{}) error {

		//strict equal must first pass the type equal strategy
		tfn := TypeEqualStrategy()
		err := tfn(example, actual)
		if err != nil {
			return err
		}

		t := newTest()
		assert.Equal(t, example, actual)
		return t.Error()
	}
}

//Actual value only needs to be of the same type as the example
func TypeEqualStrategy() AssertFunc {
	return func(example, actual interface{}) error {
		t := newTest()
		assert.IsType(t, example, actual)
		return t.Error()
	}
}
