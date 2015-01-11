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
func Assert(ex, act []byte, p parser.P) error {

	example, err := p.Parse(ex)
	if err != nil {
		return err
	}

	actual, err := p.Parse(act)
	if err != nil {
		return err
	}

	return actual.Follows(example)
}
