package postman

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLString(t *testing.T) {
	u := URL{
		Raw: "a-raw-url",
	}

	assert.Equal(t, "a-raw-url", u.String())
}

func TestCreateURLFromInterface(t *testing.T) {
	cases := []struct {
		scenario      string
		i             interface{}
		expectedURL   *URL
		expectedError error
	}{
		{
			"URL from a string",
			"http://www.google.fr",
			&URL{
				Raw: "http://www.google.fr",
			},
			nil,
		},
		{
			"URL from an interface",
			map[string]interface{}{
				"raw":      "https://gurujsonrpc.appspot.com/guru",
				"protocol": "https",
				"host":     []string{"gurujsonrpc", "appspot", "com"},
				"path":     []string{"path"},
			},
			&URL{
				Raw:      "https://gurujsonrpc.appspot.com/guru",
				Protocol: "https",
				Host:     []string{"gurujsonrpc", "appspot", "com"},
				Path:     []string{"path"},
			},
			nil,
		},
		{
			"URL from an unsupported interface",
			[]string{"not-a-valid-url"},
			nil,
			errors.New("Unsupported interface type"),
		},
	}

	for _, tc := range cases {
		req, err := createURLFromInterface(tc.i)

		assert.Equal(t, tc.expectedError, err, tc.scenario)
		assert.Equal(t, tc.expectedURL, req, tc.scenario)
	}
}

func TestURLMarsalJSON(t *testing.T) {
	cases := []struct {
		scenario       string
		url            *URL
		expectedOutput string
	}{
		{
			"Successfully marshalling an URL as a string",
			&URL{
				Raw: "http://www.google.fr",
			},
			"\"http://www.google.fr\"",
		},
		{
			"Successfully marshalling an URL with variables as a struct",
			&URL{
				Raw:      "http://www.google.fr",
				Variable: "some-variables",
			},
			"{\"raw\":\"http://www.google.fr\",\"variable\":\"some-variables\"}",
		},
	}

	for _, tc := range cases {
		bytes, _ := tc.url.MarshalJSON()

		assert.Equal(t, tc.expectedOutput, string(bytes), tc.scenario)
	}
}
