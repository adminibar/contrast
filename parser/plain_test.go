package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dockpit/contrast/parser"
)

var plain_data = []byte("some text with trailing whitespace \n\t")

var plain_adata = []byte("some text with trailing whitespace \n\t")
var plain_bdata = []byte("some text with trailing  \n\t")

func TestPlainParsing(t *testing.T) {
	p := parser.NewPlain(nil)

	table, err := p.Parse(plain_data)

	assert.NoError(t, err)
	assert.Equal(t, "some text with trailing whitespace", table.Get(".0").Value())
}

func TestPlainAtLeast_Equal(t *testing.T) {
	p := parser.NewPlain(nil)

	//when eqaul
	t1, err := p.Parse([]byte("foobar\t\n"))
	t2, err := p.Parse([]byte(`foobar`))

	assert.NoError(t, err)
	assert.NoError(t, t1.Equals(t2))
}
