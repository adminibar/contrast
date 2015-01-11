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

func TestAssert_BasicJSON_Pass(t *testing.T) {
	ats := []*assert.Archetype{}
	p := parser.NewJSON(ats)
	err := contrast.Assert([]byte(`{"foo": "bar"}`), []byte(`{"foo": "bar"}`), p)
	tassert.NoError(t, err)
}

func TestAssert_BasicJSON_Fail(t *testing.T) {
	ats := []*assert.Archetype{}
	p := parser.NewJSON(ats)
	err := contrast.Assert([]byte(`{"foo": "bar"}`), []byte(`{"foo": "rab"}`), p)
	tassert.Error(t, err)
}

func TestAssert_NestedAchetypeJSON_Pass(t *testing.T) {
	ats := []*assert.Archetype{
		&assert.Archetype{float64(42), ""},
	}

	p := parser.NewJSON(ats)
	err := contrast.Assert([]byte(`{"foo": 42}`), []byte(`{"foo": 1}`), p)
	tassert.NoError(t, err)
}
