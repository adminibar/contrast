package contrast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dockpit/contrast"
	"github.com/dockpit/contrast/parser"
)

func TestAssertJSON(t *testing.T) {
	p := parser.NewJSON()
	err := contrast.Assert([]byte("{}"), []byte("{}"), p)
	assert.NoError(t, err)
}
