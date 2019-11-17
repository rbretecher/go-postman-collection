package postman

type Body struct {
	Mode       string      `json:"mode"`
	Raw        string      `json:"raw,omitempty"`
	URLEncoded interface{} `json:"urlencoded,omitempty"`
	FormData   interface{} `json:"formdata,omitempty"`
	File       interface{} `json:"file,omitempty"`
	GraphQL    interface{} `json:"graphql,omitempty"`
	Disabled   bool        `json:"disabled,omitempty"`
}
