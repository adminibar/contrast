package assert_test

import (
	"strings"
	"testing"

	"github.com/dockpit/contrast/assert"
	tassert "github.com/stretchr/testify/assert"
)

var at_default = `[{"value": "foobar"}]`
var at_specific = `[{"value": "foobar", "strategy": "strict"}]`

func TestStringDefaultArchetype(t *testing.T) {

	//load it
	ats, err := assert.LoadArchetypes(strings.NewReader(at_default))

	tassert.NoError(t, err)
	tassert.Len(t, ats, 1)

	//parse example
	assertFn, err := assert.Parse("foobar", ats)

	tassert.NoError(t, err)
	tassert.NoError(t, assertFn("foobar", "some other string is now ok"))
	tassert.Error(t, assertFn("foobar", 20))
}

func TestStringSpecifyArchetypeStrategy(t *testing.T) {
	ats, err := assert.LoadArchetypes(strings.NewReader(at_specific))

	tassert.NoError(t, err)
	tassert.Len(t, ats, 1)

	//parse example
	assertFn, err := assert.Parse("foobar", ats)

	tassert.NoError(t, err)
	tassert.Error(t, assertFn("foobar", "some other string is now not ok"))
	tassert.Error(t, assertFn("foobar", 20))
}
