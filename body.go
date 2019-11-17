package postman

type Body struct {
	Mode       string      `json:"mode"`
	Raw        string      `json:"raw"`
	URLEncoded interface{} `json:"urlencoded"`
	FormData   interface{} `json:"formdata"`
	File       interface{} `json:"file"`
	GraphQL    interface{} `json:"graphql"`
	Disabled   bool        `json:"disabled"`
}
