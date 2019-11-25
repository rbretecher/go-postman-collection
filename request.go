package postman

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// A Request represents an HTTP request.
type Request struct {
	URL         URL         `json:"url"`
	Auth        *Auth       `json:"auth,omitempty"`
	Proxy       interface{} `json:"proxy,omitempty"`
	Certificate interface{} `json:"certificate,omitempty"`
	Method      method      `json:"method"`
	Description interface{} `json:"description,omitempty"`
	Header      []*Header   `json:"header,omitempty"`
	Body        *Body       `json:"body,omitempty"`
}

// NewRequest creates a new request.
func NewRequest(u string, m method) *Request {
	return &Request{
		URL: URL{
			Raw: u,
		},
		Method: m,
	}
}

// A Request can be created from a map[string]interface{} or a string.
// If a string, the string is assumed to be the request URL and the method is assumed to be 'GET'.
func createRequestFromInterface(i interface{}) (*Request, error) {
	switch i.(type) {
	case string:
		return NewRequest(i.(string), Get), nil
	case map[string]interface{}:
		req := new(Request)
		err := mapstructure.Decode(i, &req)
		return req, err
	default:
		return nil, errors.New("Unsupported interface type")
	}
}
