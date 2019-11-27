package postman

import (
	"errors"
	"reflect"

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
		req, err := decodeRequest(i.(map[string]interface{}))
		return req, err
	default:
		return nil, errors.New("Unsupported interface type")
	}
}

func decodeRequest(m map[string]interface{}) (req *Request, err error) {

	config := &mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &req,
		DecodeHook: func(from reflect.Type, to reflect.Type, v interface{}) (interface{}, error) {
			if to.Name() == "URL" {
				url, err := createURLFromInterface(v)
				return url, err
			}

			return v, nil
		},
	}

	decoder, err := mapstructure.NewDecoder(config)

	if err != nil {
		return
	}

	err = decoder.Decode(m)

	return
}
