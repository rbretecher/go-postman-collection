package postman

type AuthType string

const (
	ApiKey AuthType = "apikey"
	AWSV4  AuthType = "awsv4"
	Basic  AuthType = "basic"
	Bearer AuthType = "bearer"
	Digest AuthType = "digest"
	Hawk   AuthType = "hawk"
	NoAuth AuthType = "noauth"
	OAuth1 AuthType = "oauth1"
	Oauth2 AuthType = "oauth2"
	NTLM   AuthType = "ntlm"
)

type AuthParam struct {
	Key   string      `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Type  AuthType    `json:"type,omitempty"`
}

type Auth struct {
	Type   AuthType     `json:"type,omitempty"`
	Apikey []*AuthParam `json:"apikey,omitempty"`
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

func (a *Auth) GetParams() []*AuthParam {
	switch a.Type {
	case ApiKey:
		return a.Apikey
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
