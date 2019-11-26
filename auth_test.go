package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParams(t *testing.T) {

	const anUnimplementedType authType = "an-unimplemented-type"

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
		authType       authType
		expectedParams []*AuthParam
	}{
		{
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
			anUnimplementedType,
			nil,
		},
	}

	for _, tc := range cases {

		auth.Type = tc.authType

		assert.Equal(
			t,
			auth.GetParams(),
			tc.expectedParams,
		)
	}
}
