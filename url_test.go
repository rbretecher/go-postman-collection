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

func TestURLMarshalJSON(t *testing.T) {
	cases := []struct {
		scenario       string
		url            URL
		expectedOutput string
	}{
		{
			"Successfully marshalling an URL as a string",
			URL{
				Raw: "http://www.google.fr",
			},
			"\"http://www.google.fr\"",
		},
		{
			"Successfully marshalling an URL with variables as a struct",
			URL{
				Raw: "http://www.google.fr",
				Variables: []*Variable{
					{
						Name:  "a-variable",
						Value: "an-awesome-value",
					},
				},
			},
			"{\"raw\":\"http://www.google.fr\",\"variable\":[{\"name\":\"a-variable\",\"value\":\"an-awesome-value\"}]}",
		},
	}

	for _, tc := range cases {
		bytes, _ := tc.url.MarshalJSON()

		assert.Equal(t, tc.expectedOutput, string(bytes), tc.scenario)
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	cases := []struct {
		scenario      string
		bytes         []byte
		expectedURL   URL
		expectedError error
	}{
		{
			"Successfully unmarshalling an URL as a string",
			[]byte("\"http://www.google.fr\""),
			URL{
				Raw: "http://www.google.fr",
			},
			nil,
		},
		{
			"Successfully unmarshalling an URL with variables as a struct",
			[]byte("{\"raw\":\"http://www.google.fr\",\"variable\":[{\"name\":\"a-variable\",\"value\":\"an-awesome-value\"}]}"),
			URL{
				Raw: "http://www.google.fr",
				Variables: []*Variable{
					{
						Name:  "a-variable",
						Value: "an-awesome-value",
					},
				},
			},
			nil,
		},
		{
			"Failed to unmarshal an URL because of an unsupported type",
			[]byte("not-a-valid-url"),
			URL{},
			errors.New("Unsupported type"),
		},
	}

	for _, tc := range cases {

		var u URL
		err := u.UnmarshalJSON(tc.bytes)

		assert.Equal(t, tc.expectedURL, u, tc.scenario)
		assert.Equal(t, tc.expectedError, err, tc.scenario)
	}
}
