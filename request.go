package postman

import (
	"encoding/json"
	"errors"
	"fmt"
)

// A Request represents an HTTP request.
type Request struct {
	URL         *URL        `json:"url"`
	Auth        *Auth       `json:"auth,omitempty"`
	Proxy       interface{} `json:"proxy,omitempty"`
	Certificate interface{} `json:"certificate,omitempty"`
	Method      method      `json:"method"`
	Description interface{} `json:"description,omitempty"`
	Header      []*Header   `json:"header,omitempty"`
	Body        *Body       `json:"body,omitempty"`
}

// mRequest is used for marshalling/unmarshalling.
type mRequest Request

// MarshalJSON returns the JSON encoding of a Request.
// If the Request only contains an URL with the Get HTTP method, it is returned as a string.
func (r Request) MarshalJSON() ([]byte, error) {
	if r.Auth == nil && r.Proxy == nil && r.Certificate == nil && r.Description == nil && r.Header == nil && r.Body == nil && r.Method == Get {
		return []byte(fmt.Sprintf("\"%s\"", r.URL)), nil
	}

	return json.Marshal(mRequest{
		URL:         r.URL,
		Auth:        r.Auth,
		Proxy:       r.Proxy,
		Certificate: r.Certificate,
		Method:      r.Method,
		Description: r.Description,
		Header:      r.Header,
		Body:        r.Body,
	})
}

// UnmarshalJSON parses the JSON-encoded data and create a Request from it.
// A Request can be created from an object or a string.
// If a string, the string is assumed to be the request URL and the method is assumed to be 'GET'.
func (r *Request) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' {
		r.Method = Get
		r.URL = &URL{
			Raw: string(string(b[1 : len(b)-1])),
		}
	} else if b[0] == '{' {
		tmp := (*mRequest)(r)
		err = json.Unmarshal(b, &tmp)
	} else {
		err = errors.New("Unsupported type")
	}

	return
}
