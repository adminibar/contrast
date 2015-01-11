package parser

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/dockpit/contrast/assert"
)

func jsonParsingError(err error) error {
	return fmt.Errorf("JSON Parsing Error: %s", strings.Replace(err.Error(), "json: ", "", 1))
}

// A table of values that are mapped
// using json paths (e.g 0.name.full_name)
type JSONT struct {
	values     map[string]E
	archetypes []*assert.Archetype
}

func newJSONT(ats []*assert.Archetype) *JSONT {
	if ats == nil {
		ats = []*assert.Archetype{}
	}

	return &JSONT{
		values:     map[string]E{},
		archetypes: ats,
	}
}

func (t *JSONT) All() map[string]E {
	return t.values
}

func (t *JSONT) Set(key string, e E) {
	t.values[key] = e
}

func (t *JSONT) Get(key string) E {
	return t.values[key]
}

func (t *JSONT) Equals(ex T) error {
	ats := []*assert.Archetype{}

	//does this table has all the paths of
	//the example
	for path, example := range ex.All() {
		actual := t.Get(path)
		if actual == nil {
			return fmt.Errorf("Missing Value at path '%s' that example does have", path)
		}

		assert, err := example.ToAssert(ats)
		if err != nil {
			return err
		}

		err = assert(actual)
		if err != nil {
			return err
		}
	}

	//does it have any extra paths
	for path, _ := range t.All() {
		example := ex.Get(path)
		if example == nil {
			return fmt.Errorf("Extra Value at path '%s' that example doesn't have", path)
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
		t.Set(path, NewElement(e))
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

//if ats is nill, an empty list of archetypes is used
func (p *JSON) Parse(data []byte, ats []*assert.Archetype) (T, error) {
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
	t := newJSONT(ats)
	if len(l) > 0 {
		err = p.walk(l, t, "")
	} else {
		err = p.walk(o, t, "")
	}

	return t, err
}
