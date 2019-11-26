package postman

type authType string

const (
	// APIKey stands for API Key Authentication.
	APIKey authType = "apikey"
	// AWSV4 is Amazon AWS Authentication.
	AWSV4 authType = "awsv4"
	// Basic Authentication.
	Basic authType = "basic"
	// Bearer Token Authentication.
	Bearer authType = "bearer"
	// Digest Authentication.
	Digest authType = "digest"
	// Hawk Authentication.
	Hawk authType = "hawk"
	// NoAuth Authentication.
	NoAuth authType = "noauth"
	// OAuth1 Authentication.
	OAuth1 authType = "oauth1"
	// Oauth2 Authentication.
	Oauth2 authType = "oauth2"
	// NTLM Authentication.
	NTLM authType = "ntlm"
)

// AuthParam represents an attribute for any authentication method provided by Postman.
// For example "username" and "password" are set as auth attributes for Basic Authentication method.
type AuthParam struct {
	Key   string      `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Type  string      `json:"type,omitempty"`
}

// Auth contains the authentication method used and its associated parameters.
type Auth struct {
	Type   authType     `json:"type,omitempty"`
	APIKey []*AuthParam `json:"apikey,omitempty"`
	AWSV4  []*AuthParam `json:"awsv4,omitempty"`
	Basic  []*AuthParam `json:"basic,omitempty"`
	Bearer []*AuthParam `json:"bearer,omitempty"`
	Digest []*AuthParam `json:"digest,omitempty"`
	Hawk   []*AuthParam `json:"hawk,omitempty"`
	NoAuth []*AuthParam `json:"noauth,omitempty"`
	OAuth1 []*AuthParam `json:"oauth1,omitempty"`
	OAuth2 []*AuthParam `json:"oauth2,omitempty"`
	NTLM   []*AuthParam `json:"ntlm,omitempty"`
}

// GetParams returns the parameters related to the authentication method in use.
func (a Auth) GetParams() []*AuthParam {
	switch a.Type {
	case APIKey:
		return a.APIKey
	case AWSV4:
		return a.AWSV4
	case Basic:
		return a.Basic
	case Bearer:
		return a.Bearer
	case Digest:
		return a.Digest
	case Hawk:
		return a.Hawk
	case NoAuth:
		return a.NoAuth
	case OAuth1:
		return a.OAuth1
	case Oauth2:
		return a.OAuth2
	case NTLM:
		return a.NTLM
	}

	return nil
}
