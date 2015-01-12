package parser

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/dockpit/assert/strategy"
)

func jsonParsingError(err error) error {
	return fmt.Errorf("JSON Parsing Error: %s", strings.Replace(err.Error(), "json: ", "", 1))
}

// A table of values that are mapped
// using json paths (e.g 0.name.full_name)
type JSONT struct {
	values     map[string]E
	archetypes []*strategy.Archetype
}

func newJSONT(ats []*strategy.Archetype) *JSONT {
	if ats == nil {
		ats = []*strategy.Archetype{}
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

func (t *JSONT) Follows(ex T) error {

	//@todo, order ex.All() for consistent errors

	//does this table has all the paths of
	//the example
	for path, example := range ex.All() {
		actual := t.Get(path)
		if actual == nil {
			return fmt.Errorf("Missing Value at path '%s' that example does have", path)
		}

		//object/lists don't have to be checked, individual elements existence is ok
		switch actual.Value().(type) {
		case []interface{}, []map[string]interface{}, map[string]interface{}:
			continue
		}

		assert, err := example.ToAssert(t.archetypes)
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
type JSON struct {
	archetypes []*strategy.Archetype
}

func NewJSON(ats []*strategy.Archetype) *JSON {
	return &JSON{ats}
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
	t := newJSONT(p.archetypes)
	if len(l) > 0 {
		err = p.walk(l, t, "")
	} else {
		err = p.walk(o, t, "")
	}

	return t, err
}
