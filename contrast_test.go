package contrast_test

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"

	"github.com/dockpit/contrast"
	"github.com/dockpit/contrast/assert"
	"github.com/dockpit/contrast/parser"
)

func TestContentTypeToParser(t *testing.T) {
	p := contrast.Parser("application/json", nil)
	tassert.IsType(t, &parser.JSON{}, p)

	p = contrast.Parser("text/javascript", nil)
	tassert.IsType(t, &parser.Plain{}, p)
}

func TestAssertJSON(t *testing.T) {
	ats := []*assert.Archetype{}
	p := parser.NewJSON(ats)
	err := contrast.Assert([]byte("{}"), []byte("{}"), p)
	tassert.NoError(t, err)
}
