package postman

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescriptionMarshalJSON(t *testing.T) {
	cases := []struct {
		scenario       string
		description    Description
		expectedOutput string
	}{
		{
			"Successfully marshalling a Description as an object",
			Description{
				Content: "My awesome collection",
				Type:    "text/plain",
				Version: "v1",
			},
			`{"content":"My awesome collection","type":"text/plain","version":"v1"}`,
		},
		{
			"Successfully marshalling a Description as a string",
			Description{
				Content: "My awesome collection",
			},
			`"My awesome collection"`,
		},
	}

	for _, tc := range cases {
		bytes, _ := tc.description.MarshalJSON()

		assert.Equal(t, tc.expectedOutput, string(bytes), tc.scenario)
	}
}

func TestDescriptionUnmarshalJSON(t *testing.T) {
	cases := []struct {
		scenario            string
		bytes               []byte
		expectedDescription Description
		expectedError       error
	}{
		{
			"Successfully unmarshalling a Description from a string",
			[]byte(`"My awesome collection"`),
			Description{Content: "My awesome collection"},
			nil,
		},
		{
			"Successfully unmarshalling a Description from an empty slice of bytes",
			make([]byte, 0),
			Description{},
			nil,
		},
		{
			"Successfully unmarshalling a Description",
			[]byte(`{"content":"My awesome collection","type":"text/plain","version":"v1"}`),
			Description{
				Content: "My awesome collection",
				Type:    "text/plain",
				Version: "v1",
			},
			nil,
		},
		{
			"Failed to unmarshal a Description because of an unsupported type",
			[]byte(`not-a-valid-description`),
			Description{},
			errors.New("unsupported type for description"),
		},
	}

	for _, tc := range cases {

		var d Description
		err := d.UnmarshalJSON(tc.bytes)

		assert.Equal(t, tc.expectedDescription, d, tc.scenario)
		assert.Equal(t, tc.expectedError, err, tc.scenario)
	}
}
