package postman

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Header represents an HTTP Header.
type Header struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Disabled    bool   `json:"disabled,omitempty"`
	Description string `json:"description,omitempty"`
}

// HeaderList contains a list of headers.
type HeaderList struct {
	Headers []*Header
}

// MarshalJSON returns the JSON encoding of a HeaderList.
func (hl HeaderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(hl.Headers)
}

// UnmarshalJSON parses the JSON-encoded data and create a HeaderList from it.
// A HeaderList can be created from an array or a string.
func (hl *HeaderList) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	} else if len(b) >= 2 && b[0] == '"' && b[len(b)-1] == '"' {
		headersString := string(b[1 : len(b)-1])
		for _, header := range strings.Split(headersString, "\n") {
			if strings.TrimSpace(header) == "" {
				continue
			}

			headerParts := strings.Split(header, ":")

			if len(headerParts) != 2 {
				return fmt.Errorf("invalid header, missing key or value: %s", header)
			}

			hl.Headers = append(hl.Headers, &Header{
				Key:   strings.TrimSpace(headerParts[0]),
				Value: strings.TrimSpace(string(headerParts[1])),
			})
		}
	} else if len(b) >= 2 && b[0] == '[' && b[len(b)-1] == ']' {
		err = json.Unmarshal(b, &hl.Headers)
	} else {
		err = errors.New("unsupported type for header list")
	}

	return
}
