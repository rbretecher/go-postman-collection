package postman

// A Variable allows you to store and reuse values in your requests and scripts.
type Variable struct {
	ID          string `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	System      bool   `json:"system,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
}

// CreateVariable creates a new Variable of type string.
func CreateVariable(name string, value string) *Variable {
	return &Variable{
		Name:  name,
		Value: value,
		Type:  "string",
	}
}
