package postman

type Request struct {
	URL         interface{} `json:"url"`
	Auth        interface{} `json:"auth"`
	Proxy       interface{} `json:"proxy"`
	Certificate interface{} `json:"certificate"`
	Method      string      `json:"method"`
	Description interface{} `json:"description"`
	Header      interface{} `json:"header"`
	Body        interface{} `json:"body"`
}

func NewRequest(URL string, method string) *Request {
	return &Request{
		URL:    URL,
		Method: method,
	}
}
