package assert

import (
	"github.com/dockpit/assert/parser"
	"github.com/dockpit/assert/strategy"
)

var DefaultParserMIME = "plain/text"

//Return a parser that is appropriate for the given mime type, if
//non is found return parser for DefaultParserMIME
func Parser(mime string, ats []*strategy.Archetype) parser.P {
	switch mime {
	case "application/json":
		return parser.NewJSON(ats)
	}

	return parser.NewPlain(ats)
}

//Assert whether the given data flollows the example bytes
//using the given parser
func Follows(ex, act []byte, p parser.P) error {

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
