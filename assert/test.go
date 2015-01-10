package assert

import (
	"fmt"
)

type test struct {
	Errors []error
}

func newTest() *test {
	return &test{[]error{}}
}

func (t *test) Error() error {
	if len(t.Errors) == 0 {
		return nil
	}

	str := ""
	for _, e := range t.Errors {
		str = fmt.Sprintf("%s\n%s", str, e)
	}
	return fmt.Errorf(str)
}

func (t *test) Errorf(format string, args ...interface{}) {
	t.Errors = append(t.Errors, fmt.Errorf(format, args...))
}
