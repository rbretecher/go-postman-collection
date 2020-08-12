package postman

// These constants represent the available raw languages.
const (
	HTML       string = "html"
	Javascript string = "javascript"
	JSON       string = "json"
	Text       string = "text"
	XML        string = "xml"
)

// Body represents the data usually contained in the request body.
type Body struct {
	Mode       string       `json:"mode"`
	Raw        string       `json:"raw,omitempty"`
	URLEncoded interface{}  `json:"urlencoded,omitempty"`
	FormData   interface{}  `json:"formdata,omitempty"`
	File       interface{}  `json:"file,omitempty"`
	GraphQL    interface{}  `json:"graphql,omitempty"`
	Disabled   bool         `json:"disabled,omitempty"`
	Options    *BodyOptions `json:"options,omitempty"`
}

// BodyOptions holds body options.
type BodyOptions struct {
	Raw BodyOptionsRaw `json:"raw,omitempty"`
}

// BodyOptionsRaw represents the acutal language to use in postman. (See possible options in the cost above)
type BodyOptionsRaw struct {
	Language string `json:"language,omitempty"`
}
