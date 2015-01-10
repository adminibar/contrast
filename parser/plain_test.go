package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dockpit/contrast/parser"
)

var data = []byte("some text with trailing whitespace \n\t")

func TestPlainParsing(t *testing.T) {
	p := parser.NewPlain()

	table, err := p.Parse(data)

	assert.NoError(t, err)
	assert.Equal(t, "some text with trailing whitespace", table.Get(".0").Value())
}
