package postman

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// URL is a struct that contains an URL in a "broken-down way".
// Raw contains the complete URL.
type URL struct {
	Raw      string      `json:"raw"`
	Protocol string      `json:"protocol,omitempty"`
	Host     []string    `json:"host,omitempty"`
	Path     []string    `json:"path,omitempty"`
	Port     string      `json:"port,omitempty"`
	Query    interface{} `json:"query,omitempty"`
	Hash     string      `json:"hash,omitempty"`
	Variable interface{} `json:"variable,omitempty"`
}

// Used to Marshall the URL without calling the URL.MarshalJSON function.
type marshalledURL URL

// String returns the raw version of the URL.
func (u URL) String() string {
	return u.Raw
}

// An URL can be created from a map[string]interface{} or a string.
// If a string, the string is assumed to be the Raw attribute.
func createURLFromInterface(i interface{}) (*URL, error) {
	switch i.(type) {
	case string:
		return &URL{Raw: i.(string)}, nil
	case map[string]interface{}:
		url := new(URL)
		err := mapstructure.Decode(i, &url)
		return url, err
	default:
		return nil, errors.New("Unsupported interface type")
	}
}

// MarshalJSON marshalls the URL as a string if it does not contain any variable.
// In case it contains any variable, it gets marshalled as a struct.
func (u *URL) MarshalJSON() ([]byte, error) {

	if u.Variable == nil {
		return []byte(fmt.Sprintf("\"%s\"", u.Raw)), nil
	}

	return json.Marshal(marshalledURL{
		Raw:      u.Raw,
		Protocol: u.Protocol,
		Host:     u.Host,
		Path:     u.Path,
		Port:     u.Port,
		Query:    u.Query,
		Hash:     u.Hash,
		Variable: u.Variable,
	})
}
