package postman

import (
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
