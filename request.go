package postman

type Request struct {
	URL         interface{} `json:"url"`
	Auth        interface{} `json:"auth,omitempty"`
	Proxy       interface{} `json:"proxy,omitempty"`
	Certificate interface{} `json:"certificate,omitempty"`
	Method      string      `json:"method"`
	Description interface{} `json:"description,omitempty"`
	Header      []*Header   `json:"header,omitempty"`
	Body        *Body       `json:"body,omitempty"`
}

func NewRequest(URL string, method Method) *Request {
	return &Request{
		URL:    URL,
		Method: method.String(),
	}
}
