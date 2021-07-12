package postman

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Description contains collection description.
type Description struct {
	Content string `json:"content,omitempty"`
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
}

// mDescription is used for marshalling/unmarshalling.
type mDescription Description

// MarshalJSON returns the JSON encoding of a Description.
// If the Description only has a content, it is returned as a string.
func (d Description) MarshalJSON() ([]byte, error) {
	if d.Type == "" && d.Version == "" {
		return []byte(fmt.Sprintf("\"%s\"", d.Content)), nil
	}

	return json.Marshal(mDescription{
		Content: d.Content,
		Type:    d.Type,
		Version: d.Version,
	})
}

// UnmarshalJSON parses the JSON-encoded data and create a Description from it.
// A Description can be created from an object or a string.
func (d *Description) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	} else if len(b) >= 2 && b[0] == '"' && b[len(b)-1] == '"' {
		d.Content = string(string(b[1 : len(b)-1]))
	} else if len(b) >= 2 && b[0] == '{' && b[len(b)-1] == '}' {
		tmp := (*mDescription)(d)
		err = json.Unmarshal(b, &tmp)
	} else {
		err = errors.New("unsupported type for description")
	}

	return
}
