package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dockpit/contrast/parser"
)

var odata = []byte(`{"foo": "bar", "things": [{"a": [1,2,4]}, {"key": "value"}]}`)
var ldata = []byte(`[` + string(odata) + `]`)

var adata = []byte(`{"foo": "bar", "lorum": "ipsum"}`)
var bdata = []byte(`{"foo": "bar", "lorum": "ipsum"}`)

func TestJSONObjectParsing(t *testing.T) {

	p := parser.NewJSON()

	ta, err := p.Parse(odata, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, ta.Get(".things.0.a.1").Value())
}

func TestJSONListParsing(t *testing.T) {

	p := parser.NewJSON()

	ta, err := p.Parse(ldata, nil)
	assert.NoError(t, err)

	assert.Equal(t, "bar", ta.Get(".0.foo").Value())
	assert.Equal(t, 2, ta.Get(".0.things.0.a.1").Value())
}

func TestJSONAtLeast_Equal(t *testing.T) {
	p := parser.NewJSON()

	//when eqaul
	t1, err := p.Parse([]byte(`{"foo": "bar", "lorum": "ipsum"}`), nil)
	t2, err := p.Parse([]byte(`{"foo": "bar", "lorum": "ipsum"}`), nil)

	assert.NoError(t, err)
	assert.NoError(t, t1.AtLeast(t2))
}

func TestJSONAtLeast_Less(t *testing.T) {
	p := parser.NewJSON()

	//when eqaul
	t1, err := p.Parse([]byte(`{"foo": "bar", "lorum": "ipsum"}`), nil)
	t2, err := p.Parse([]byte(`{"foo": "bar"}`), nil)

	assert.NoError(t, err)
	assert.NoError(t, t1.AtLeast(t2))
}

func TestJSONAtLeast_More(t *testing.T) {
	p := parser.NewJSON()

	//when eqaul
	t1, err := p.Parse([]byte(`{"foo": "bar"}`), nil)
	t2, err := p.Parse([]byte(`{"foo": "bar", "lorum": "ipsum"}`), nil)

	assert.NoError(t, err)
	assert.Error(t, t1.AtLeast(t2))
}
