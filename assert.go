package contrast

import (
	"github.com/dockpit/contrast/parser"
)

//Assert the given data against the example bytes using
//the content type specific parser
func Assert(ex, data []byte, p parser.P) error {

	// t1 := p.parse(ex)

	// t2 := p.parse(data)

	// d := t1.Diff(t2)
	// c := t1.AtMost(t2)
	// c := t1.Exact(t2)

	return nil
}
