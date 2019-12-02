package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLString(t *testing.T) {
	u := URL{
		Raw: "a-raw-url",
	}

	assert.Equal(t, "a-raw-url", u.String())
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
