package postman

// URL is a struct that contains an URL in a "broken-down way".
// Raw contains the complete URL.
type URL struct {
	Raw      string      `json:"raw"`
	Protocol string      `json:"protocol,omitempty"`
	Host     []string    `json:"host,omitempty"`
	Path     []string    `json:"path,omitempty"`
	Port     string      `json:"port,omitempty"`
	Query    interface{} `json:"query,omitempty"`
	Hash     string      `json:"hash,omitempty"`
	Variable interface{} `json:"variable,omitempty"`
}

// String returns the raw version of the URL.
func (u URL) String() string {
	return u.Raw
}
