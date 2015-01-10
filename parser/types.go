package parser

type T interface {
	All() map[string]interface{}
	Set(string, interface{})
	Get(string) interface{}
	AtLeast(T) error
}

type P interface {
	Parse([]byte) (T, error)
}
