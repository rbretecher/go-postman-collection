package postman

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
