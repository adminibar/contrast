package parser

type JSON struct{}

func NewJSON() *JSON {
	return &JSON{}
}

func (p *JSON) Parse([]byte) (error, map[string]string) {
	return nil, map[string]string{}
}
