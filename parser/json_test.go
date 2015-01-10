package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dockpit/contrast/parser"
)

var odata = []byte(`{"foo": "bar", "things": [{"a": [1,2,4]}, {"key": "value"}]}`)
var ldata = []byte(`[` + string(odata) + `]`)

func TestJSONObjectParsing(t *testing.T) {

	p := parser.NewJSON()

	ta, err := p.Parse(odata)
	assert.NoError(t, err)
	assert.Equal(t, 2, ta.Get(".things.0.a.1"))
}

func TestJSONListParsing(t *testing.T) {

	p := parser.NewJSON()

	ta, err := p.Parse(ldata)
	assert.NoError(t, err)

	assert.Equal(t, "bar", ta.Get(".0.foo"))
	assert.Equal(t, 2, ta.Get(".0.things.0.a.1"))
}
