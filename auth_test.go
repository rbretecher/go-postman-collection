package postman

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParams(t *testing.T) {
	auth := Auth{
		Type: APIKey,
		APIKey: []*AuthParam{
			{
				Type:  "string",
				Key:   "apikey-key",
				Value: "apikey-value",
			},
		},
		AWSV4: []*AuthParam{
			{
				Type:  "string",
				Key:   "awsv4-key",
				Value: "awsv4-value",
			},
		},
		Basic: []*AuthParam{
			{
				Type:  "string",
				Key:   "basic-key",
				Value: "basic-value",
			},
		},
		Bearer: []*AuthParam{
			{
				Type:  "string",
				Key:   "bearer-key",
				Value: "bearer-value",
			},
		},
		Digest: []*AuthParam{
			{
				Type:  "string",
				Key:   "digest-key",
				Value: "digest-value",
			},
		},
		Hawk: []*AuthParam{
			{
				Type:  "string",
				Key:   "hawk-key",
				Value: "hawk-value",
			},
		},
		NoAuth: []*AuthParam{
			{
				Type:  "string",
				Key:   "noauth-key",
				Value: "noauth-value",
			},
		},
		OAuth1: []*AuthParam{
			{
				Type:  "string",
				Key:   "oauth1-key",
				Value: "oauth1-value",
			},
		},
		OAuth2: []*AuthParam{
			{
				Type:  "string",
				Key:   "oauth2-key",
				Value: "oauth2-value",
			},
		},
		NTLM: []*AuthParam{
			{
				Type:  "string",
				Key:   "ntlm-key",
				Value: "ntlm-value",
			},
		},
	}

	cases := []struct {
		scenario       string
		authType       authType
		expectedParams []*AuthParam
	}{
		{
			"GetParams for ApiKey",
			APIKey,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "apikey-key",
					Value: "apikey-value",
				},
			},
		},
		{
			"GetParams for AWSV4",
			AWSV4,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "awsv4-key",
					Value: "awsv4-value",
				},
			},
		},
		{
			"GetParams for Basic",
			Basic,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "basic-key",
					Value: "basic-value",
				},
			},
		},
		{
			"GetParams for Bearer",
			Bearer,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "bearer-key",
					Value: "bearer-value",
				},
			},
		},
		{
			"GetParams for Digest",
			Digest,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "digest-key",
					Value: "digest-value",
				},
			},
		},
		{
			"GetParams for Hawk",
			Hawk,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "hawk-key",
					Value: "hawk-value",
				},
			},
		},
		{
			"GetParams for NoAuth",
			NoAuth,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "noauth-key",
					Value: "noauth-value",
				},
			},
		},
		{
			"GetParams for OAuth1",
			OAuth1,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "oauth1-key",
					Value: "oauth1-value",
				},
			},
		},
		{
			"GetParams for Oauth2",
			Oauth2,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "oauth2-key",
					Value: "oauth2-value",
				},
			},
		},
		{
			"GetParams for NTLM",
			NTLM,
			[]*AuthParam{
				{
					Type:  "string",
					Key:   "ntlm-key",
					Value: "ntlm-value",
				},
			},
		},
		{
			"GetParams for an unimplemented authentication method",
			"an-unimplemented-authentication-method",
			nil,
		},
	}

	for _, tc := range cases {
		auth.Type = tc.authType

		assert.Equal(
			t,
			tc.expectedParams,
			auth.GetParams(),
			tc.scenario,
		)
	}
}

func TestAuthUnmarshalJSON(t *testing.T) {
	cases := []struct {
		scenario      string
		bytes         []byte
		expectedAuth  *Auth
		expectedError error
	}{
		{
			"Successfully unmarshalling a basic Auth v2.0.0",
			[]byte("{\"type\":\"basic\",\"basic\":{\"a-key\":\"a-value\"}}"),
			&Auth{
				Type: Basic,
				Basic: []*AuthParam{
					{
						Key:   "a-key",
						Value: "a-value",
					},
				},
			},
			nil,
		},
		{
			"Successfully unmarshalling a basic Auth v2.1.0",
			[]byte("{\"type\":\"basic\",\"basic\":[{\"key\":\"a-key\",\"value\":\"a-value\"}]}"),
			&Auth{
				Type: Basic,
				Basic: []*AuthParam{
					{
						Key:   "a-key",
						Value: "a-value",
					},
				},
			},
			nil,
		},

		{
			"Failed to unmarshal apiKey auth because of an unsupported format",
			[]byte("{\"type\":\"apikey\",\"apikey\":\"invalid-auth-param\"}"),
			&Auth{
				Type: APIKey,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal awsv4 auth because of an unsupported format",
			[]byte("{\"type\":\"awsv4\",\"awsv4\":\"invalid-auth-param\"}"),
			&Auth{
				Type: AWSV4,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal basic auth because of an unsupported format",
			[]byte("{\"type\":\"basic\",\"basic\":\"invalid-auth-param\"}"),
			&Auth{
				Type: Basic,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal bearer auth because of an unsupported format",
			[]byte("{\"type\":\"bearer\",\"bearer\":\"invalid-auth-param\"}"),
			&Auth{
				Type: Bearer,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal digest auth because of an unsupported format",
			[]byte("{\"type\":\"digest\",\"digest\":\"invalid-auth-param\"}"),
			&Auth{
				Type: Digest,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal hawk auth because of an unsupported format",
			[]byte("{\"type\":\"hawk\",\"hawk\":\"invalid-auth-param\"}"),
			&Auth{
				Type: Hawk,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal noauth auth because of an unsupported format",
			[]byte("{\"type\":\"noauth\",\"noauth\":\"invalid-auth-param\"}"),
			&Auth{
				Type: NoAuth,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal oauth1 auth because of an unsupported format",
			[]byte("{\"type\":\"oauth1\",\"oauth1\":\"invalid-auth-param\"}"),
			&Auth{
				Type: OAuth1,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal oauth2 auth because of an unsupported format",
			[]byte("{\"type\":\"oauth2\",\"oauth2\":\"invalid-auth-param\"}"),
			&Auth{
				Type: Oauth2,
			},
			errors.New("Unsupported type"),
		},
		{
			"Failed to unmarshal ntlm auth because of an unsupported format",
			[]byte("{\"type\":\"ntlm\",\"ntlm\":\"invalid-auth-param\"}"),
			&Auth{
				Type: NTLM,
			},
			errors.New("Unsupported type"),
		},
	}

	for _, tc := range cases {

		a := new(Auth)
		err := a.UnmarshalJSON(tc.bytes)

		assert.Equal(t, tc.expectedAuth, a, tc.scenario)
		assert.Equal(t, tc.expectedError, err, tc.scenario)
	}
}

func TestAuthMarshalJSON(t *testing.T) {
	cases := []struct {
		scenario       string
		auth           *Auth
		expectedOutput string
	}{
		{
			"Successfully marshalling an Auth v2.0.0",
			&Auth{
				version: V200,
				Type:    Basic,
				Basic: []*AuthParam{
					{
						Key:   "a-key",
						Value: "a-value",
					},
				},
			},
			"{\"type\":\"basic\",\"basic\":{\"a-key\":\"a-value\"}}",
		},
		{
			"Successfully marshalling an Auth v2.1.0",
			&Auth{
				version: V210,
				Type:    Basic,
				Basic: []*AuthParam{
					{
						Key:   "a-key",
						Value: "a-value",
					},
				},
			},
			"{\"type\":\"basic\",\"basic\":[{\"key\":\"a-key\",\"value\":\"a-value\"}]}",
		},
	}

	for _, tc := range cases {
		bytes, _ := tc.auth.MarshalJSON()

		assert.Equal(t, tc.expectedOutput, string(bytes), tc.scenario)
	}
}
