package postman

import (
	"encoding/json"
	"errors"
)

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
type Auth authV210

type authV210 struct {
	version version
	Type    authType     `json:"type,omitempty"`
	APIKey  []*AuthParam `json:"apikey,omitempty"`
	AWSV4   []*AuthParam `json:"awsv4,omitempty"`
	Basic   []*AuthParam `json:"basic,omitempty"`
	Bearer  []*AuthParam `json:"bearer,omitempty"`
	Digest  []*AuthParam `json:"digest,omitempty"`
	Hawk    []*AuthParam `json:"hawk,omitempty"`
	NoAuth  []*AuthParam `json:"noauth,omitempty"`
	OAuth1  []*AuthParam `json:"oauth1,omitempty"`
	OAuth2  []*AuthParam `json:"oauth2,omitempty"`
	NTLM    []*AuthParam `json:"ntlm,omitempty"`
}

type authV200 struct {
	Type   authType               `json:"type,omitempty"`
	APIKey map[string]interface{} `json:"apikey,omitempty"`
	AWSV4  map[string]interface{} `json:"awsv4,omitempty"`
	Basic  map[string]interface{} `json:"basic,omitempty"`
	Bearer map[string]interface{} `json:"bearer,omitempty"`
	Digest map[string]interface{} `json:"digest,omitempty"`
	Hawk   map[string]interface{} `json:"hawk,omitempty"`
	NoAuth map[string]interface{} `json:"noauth,omitempty"`
	OAuth1 map[string]interface{} `json:"oauth1,omitempty"`
	OAuth2 map[string]interface{} `json:"oauth2,omitempty"`
	NTLM   map[string]interface{} `json:"ntlm,omitempty"`
}

// mAuth is used for marshalling/unmarshalling.
type mAuth struct {
	Type   authType        `json:"type,omitempty"`
	APIKey json.RawMessage `json:"apikey,omitempty"`
	AWSV4  json.RawMessage `json:"awsv4,omitempty"`
	Basic  json.RawMessage `json:"basic,omitempty"`
	Bearer json.RawMessage `json:"bearer,omitempty"`
	Digest json.RawMessage `json:"digest,omitempty"`
	Hawk   json.RawMessage `json:"hawk,omitempty"`
	NoAuth json.RawMessage `json:"noauth,omitempty"`
	OAuth1 json.RawMessage `json:"oauth1,omitempty"`
	OAuth2 json.RawMessage `json:"oauth2,omitempty"`
	NTLM   json.RawMessage `json:"ntlm,omitempty"`
}

func (a *Auth) setVersion(v version) {
	a.version = v
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

// UnmarshalJSON parses the JSON-encoded data and create an Auth from it.
// Depending on the Postman Collection version, an auth property can either be an array or an object.
//    - v2.1.0 : Array
//    - v2.0.0 : Object
func (a *Auth) UnmarshalJSON(b []byte) (err error) {
	var tmp mAuth
	err = json.Unmarshal(b, &tmp)

	a.Type = tmp.Type

	if a.APIKey, err = unmarshalAuthParam(tmp.APIKey); err != nil {
		return
	}
	if a.AWSV4, err = unmarshalAuthParam(tmp.AWSV4); err != nil {
		return
	}
	if a.Basic, err = unmarshalAuthParam(tmp.Basic); err != nil {
		return
	}
	if a.Bearer, err = unmarshalAuthParam(tmp.Bearer); err != nil {
		return
	}
	if a.Digest, err = unmarshalAuthParam(tmp.Digest); err != nil {
		return
	}
	if a.Hawk, err = unmarshalAuthParam(tmp.Hawk); err != nil {
		return
	}
	if a.NoAuth, err = unmarshalAuthParam(tmp.NoAuth); err != nil {
		return
	}
	if a.OAuth1, err = unmarshalAuthParam(tmp.OAuth1); err != nil {
		return
	}
	if a.OAuth2, err = unmarshalAuthParam(tmp.OAuth2); err != nil {
		return
	}
	if a.NTLM, err = unmarshalAuthParam(tmp.NTLM); err != nil {
		return
	}

	return
}

func unmarshalAuthParam(b []byte) (a []*AuthParam, err error) {
	if len(b) > 0 {
		if b[0] == '{' { // v2.0.0
			var tmp map[string]string
			json.Unmarshal(b, &tmp)
			for k, v := range tmp {
				a = append(a, &AuthParam{
					Key:   k,
					Value: v,
				})
			}
		} else if b[0] == '[' { // v2.1.0
			json.Unmarshal(b, &a)
		} else {
			err = errors.New("Unsupported type")
		}
	}

	return
}

// MarshalJSON returns the JSON encoding of a Auth.
// If the version is v2.0.0 it is returned as an object, otherwise as an array (v2.1.0).
func (a Auth) MarshalJSON() ([]byte, error) {

	if a.version == V200 {
		return json.Marshal(authV200{
			Type:   a.Type,
			APIKey: authParamsToMap(a.APIKey),
			AWSV4:  authParamsToMap(a.AWSV4),
			Basic:  authParamsToMap(a.Basic),
			Bearer: authParamsToMap(a.Bearer),
			Digest: authParamsToMap(a.Digest),
			Hawk:   authParamsToMap(a.Hawk),
			NoAuth: authParamsToMap(a.NoAuth),
			OAuth1: authParamsToMap(a.OAuth1),
			OAuth2: authParamsToMap(a.OAuth2),
			NTLM:   authParamsToMap(a.NTLM),
		})
	}

	return json.Marshal(authV210{
		Type:   a.Type,
		APIKey: a.APIKey,
		AWSV4:  a.AWSV4,
		Basic:  a.Basic,
		Bearer: a.Bearer,
		Digest: a.Digest,
		Hawk:   a.Hawk,
		NoAuth: a.NoAuth,
		OAuth1: a.OAuth1,
		OAuth2: a.OAuth2,
		NTLM:   a.NTLM,
	})
}

func authParamsToMap(authParams []*AuthParam) map[string]interface{} {
	authParamsMap := make(map[string]interface{})

	for _, authParam := range authParams {
		authParamsMap[authParam.Key] = authParam.Value
	}

	return authParamsMap
}
