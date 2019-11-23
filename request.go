package postman

type Request struct {
	URL         URL         `json:"url"`
	Auth        interface{} `json:"auth,omitempty"`
	Proxy       interface{} `json:"proxy,omitempty"`
	Certificate interface{} `json:"certificate,omitempty"`
	Method      string      `json:"method"`
	Description interface{} `json:"description,omitempty"`
	Header      []*Header   `json:"header,omitempty"`
	Body        *Body       `json:"body,omitempty"`
}

func NewRequest(u string, method Method) *Request {
	return &Request{
		URL: URL{
			Raw: u,
		},
		Method: method.String(),
	}
}
