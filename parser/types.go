package parser

type P interface {
	Parse([]byte) (error, map[string]string)
}
