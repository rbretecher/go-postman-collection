package postman

// Body represents the data usually contained in the request body.
type Body struct {
	Mode       string      `json:"mode"`
	Raw        string      `json:"raw,omitempty"`
	URLEncoded interface{} `json:"urlencoded,omitempty"`
	FormData   interface{} `json:"formdata,omitempty"`
	File       interface{} `json:"file,omitempty"`
	GraphQL    interface{} `json:"graphql,omitempty"`
	Disabled   bool        `json:"disabled,omitempty"`
	Options    Raw         `json:"options,omitempty"`
}

//Raw represents the data of options->language in postman
type Raw struct {
	Raw interface{} `json:"raw,omitempty"`
}
