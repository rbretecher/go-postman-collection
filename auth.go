package postman

type AuthType string

const (
	APIKEY AuthType = "apikey"
	AWSV4  AuthType = "awsv4"
	BASIC  AuthType = "basic"
	BEARER AuthType = "bearer"
	DIGEST AuthType = "digest"
	HAWK   AuthType = "hawk"
	NOAUTH AuthType = "noauth"
	OAUTH1 AuthType = "oauth1"
	OAUTH2 AuthType = "oauth2"
	NTLM   AuthType = "ntlm"
)

type AuthParam struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Type  AuthType    `json:"type"`
}

type Auth struct {
	Type   AuthType
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
	case APIKEY:
		return a.Apikey
	case AWSV4:
		return a.AWSV4
	case BASIC:
		return a.Basic
	case BEARER:
		return a.Bearer
	case DIGEST:
		return a.Digest
	case HAWK:
		return a.Hawk
	case NOAUTH:
		return a.NoAuth
	case OAUTH1:
		return a.OAuth1
	case OAUTH2:
		return a.OAuth2
	case NTLM:
		return a.NTLM
	}

	return nil
}
