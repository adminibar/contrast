package assert_test

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"

	"github.com/dockpit/assert"
	"github.com/dockpit/assert/parser"
	"github.com/dockpit/assert/strategy"
)

var ats = []*strategy.Archetype{
	&strategy.Archetype{float64(42), ""},
}

func TestContentTypeToParser(t *testing.T) {
	p := assert.Parser("application/json", nil)
	tassert.IsType(t, &parser.JSON{}, p)

	p = assert.Parser("text/javascript", nil)
	tassert.IsType(t, &parser.Plain{}, p)
}

func TestAssert_BasicJSON_Pass(t *testing.T) {
	p := parser.NewJSON(ats)
	err := assert.Follows([]byte(`{"foo": "bar"}`), []byte(`{"foo": "bar"}`), p)
	tassert.NoError(t, err)
}

func TestAssert_BasicJSON_Fail(t *testing.T) {
	p := parser.NewJSON(ats)
	err := assert.Follows([]byte(`{"foo": "bar"}`), []byte(`{"foo": "rab"}`), p)
	tassert.Error(t, err)
}

func TestAssert_AchetypeJSON_Pass(t *testing.T) {
	p := parser.NewJSON(ats)
	err := assert.Follows([]byte(`{"foo": 42}`), []byte(`{"foo": 1}`), p)
	tassert.NoError(t, err)
}

func TestAssert_NestedAchetypeJSON_Pass(t *testing.T) {
	p := parser.NewJSON(ats)
	err := assert.Follows([]byte(`{"foo": {"bar": [42]}}`), []byte(`{"foo": {"bar": [1]}}`), p)
	tassert.NoError(t, err)
}

func TestAssert_NestedAchetypeJSON_Fail(t *testing.T) {
	p := parser.NewJSON(ats)
	err := assert.Follows([]byte(`{"foo": {"bar": [43]}}`), []byte(`{"foo": {"bar": [1]}}`), p)
	tassert.Error(t, err)
}

func TestAssert_Plain_Fail(t *testing.T) {
	p := parser.NewPlain(ats)
	err := assert.Follows([]byte(`abcd`), []byte(`abcde`), p)
	tassert.Error(t, err)
}

func TestAssert_Plain_Success(t *testing.T) {
	p := parser.NewPlain(ats)
	err := assert.Follows([]byte("abcd \n\t"), []byte("abcd "), p)
	tassert.NoError(t, err)
}
