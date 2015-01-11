package assert

import (
	"encoding/json"
	"io"
)

//Archetype allows the developer to
//define certain content to trigger a
//different assertion strategy instead
//of the default 'strict' strategy
type Archetype struct {
	Value    interface{} `json:"value"`
	Strategy string      `json:"strategy"`
}

//returns whether a given example value fit the defined archetype
func (a *Archetype) Fits(val interface{}) bool {
	return a.Value == val
}

//returns the assert strategy for the archetype
func (a *Archetype) ToStrategy() (Strategy, error) {
	for name, st := range AvailableStrategies {
		if a.Strategy == name {
			return st, nil
		}
	}

	//default to type equal strategy if we have an
	//archetype matches an example
	return TypeEqualStrategy, nil
}

//load archetypes from json encoded reader content
func LoadArchetypes(r io.Reader) ([]*Archetype, error) {
	ats := []*Archetype{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&ats)
	return ats, err
}
