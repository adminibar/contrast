package parser

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func jsonParsingError(err error) error {
	return fmt.Errorf("JSON Parsing Error: %s", strings.Replace(err.Error(), "json: ", "", 1))
}

// A table of values that are mapped
// using json paths (e.g 0.name.full_name)
type JSONT struct {
	values map[string]interface{}
}

func newJSONT() *JSONT {
	return &JSONT{
		values: map[string]interface{}{},
	}
}

func (t *JSONT) All() map[string]interface{} {
	return t.values
}

func (t *JSONT) Set(key string, val interface{}) {
	t.values[key] = val
}

func (t *JSONT) Get(key string) interface{} {
	return t.values[key]
}

// Asserts wether this table has at least the
// content of the other (example) table
func (t *JSONT) AtLeast(ex T) error {

	for path, example := range ex.All() {
		actual := t.Get(path)

		//
		//@todo, add more advanced assertion options
		//

		if actual != example {
			return fmt.Errorf("%s != %s", actual, example)
		}
	}

	return nil
}

// For parsing byte arrays that are known to be
// JSON encoded
type JSON struct{}

func NewJSON() *JSON {
	return &JSON{}
}

func (p *JSON) walk(e interface{}, t *JSONT, path string) error {
	var err error

	//skip root
	if path != "" {
		t.Set(path, e)
	}

	switch et := e.(type) {
	case []interface{}:
		for i, e := range et {
			err = p.walk(e, t, path+"."+strconv.Itoa(i))
			if err != nil {
				return err
			}
		}
	case []map[string]interface{}:
		for i, e := range et {
			err = p.walk(e, t, path+"."+strconv.Itoa(i))
			if err != nil {
				return err
			}
		}
	case map[string]interface{}:
		for k, v := range et {
			err = p.walk(v, t, path+"."+k)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *JSON) Parse(data []byte) (T, error) {
	l := []map[string]interface{}{}
	o := map[string]interface{}{}
	err := json.Unmarshal(data, &l)
	if err != nil {

		//if it is a json object instead unmarshal as such
		if strings.Contains(err.Error(), "cannot unmarshal object") {
			err := json.Unmarshal(data, &o)
			if err != nil {
				return nil, jsonParsingError(err)
			}
		} else {
			return nil, jsonParsingError(err)
		}

	}

	//walk either the list or the object
	t := newJSONT()
	if len(l) > 0 {
		err = p.walk(l, t, "")
	} else {
		err = p.walk(o, t, "")
	}

	return t, err
}
