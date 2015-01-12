package strategy_test

import (
	"strings"
	"testing"

	"github.com/dockpit/assert/strategy"
	"github.com/stretchr/testify/assert"
)

var at_default = `[{"value": "foobar"}]`
var at_specific = `[{"value": "foobar", "strategy": "strict"}]`

func TestStringDefaultArchetype(t *testing.T) {

	//load it
	ats, err := strategy.LoadArchetypes(strings.NewReader(at_default))

	assert.NoError(t, err)
	assert.Len(t, ats, 1)

	//parse example
	assertFn, err := strategy.Parse("foobar", ats)

	assert.NoError(t, err)
	assert.NoError(t, assertFn("foobar", "some other string is now ok"))
	assert.Error(t, assertFn("foobar", 20))
}

func TestStringSpecifyArchetypeStrategy(t *testing.T) {
	ats, err := strategy.LoadArchetypes(strings.NewReader(at_specific))

	assert.NoError(t, err)
	assert.Len(t, ats, 1)

	//parse example
	assertFn, err := strategy.Parse("foobar", ats)

	assert.NoError(t, err)
	assert.Error(t, assertFn("foobar", "some other string is now not ok"))
	assert.Error(t, assertFn("foobar", 20))
}
