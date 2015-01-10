package parser

type T interface {
	Set(string, interface{})
	Get(string) interface{}
	AtLeast(T) error
}

type P interface {
	Parse([]byte) (T, error)
}
