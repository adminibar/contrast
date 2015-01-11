package contrast

import (
	"github.com/dockpit/contrast/assert"
	"github.com/dockpit/contrast/parser"
)

var DefaultParserMIME = "plain/text"

//Return a parser that is appropriate for the given mime type, if
//non is found return parser for DefaultParserMIME
func Parser(mime string, ats []*assert.Archetype) parser.P {
	switch mime {
	case "application/json":
		return parser.NewJSON(ats)
	}

	return parser.NewPlain(ats)
}

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
