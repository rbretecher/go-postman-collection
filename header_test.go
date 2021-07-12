package postman

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderListMarshalJSON(t *testing.T) {
	cases := []struct {
		scenario       string
		headerList     HeaderList
		expectedOutput string
	}{
		{
			"Successfully marshalling a HeaderList",
			HeaderList{
				Headers: []*Header{
					{
						Key:         "Content-Type",
						Value:       "application/json",
						Description: "The content type",
					},
					{
						Key:         "Authorization",
						Value:       "Bearer a-bearer-token",
						Description: "The Bearer token",
					},
				},
			},
			"[{\"key\":\"Content-Type\",\"value\":\"application/json\",\"description\":\"The content type\"},{\"key\":\"Authorization\",\"value\":\"Bearer a-bearer-token\",\"description\":\"The Bearer token\"}]",
		},
	}

	for _, tc := range cases {
		bytes, _ := tc.headerList.MarshalJSON()

		assert.Equal(t, tc.expectedOutput, string(bytes), tc.scenario)
	}
}

func TestHeaderListUnmarshalJSON(t *testing.T) {
	cases := []struct {
		scenario           string
		bytes              []byte
		expectedHeaderList HeaderList
		expectedError      error
	}{
		{
			"Successfully unmarshalling a HeaderList from a string",
			[]byte("\"Content-Type: application/json\nAuthorization: Bearer a-bearer-token\n\""),
			HeaderList{
				Headers: []*Header{
					{
						Key:   "Content-Type",
						Value: "application/json",
					},
					{
						Key:   "Authorization",
						Value: "Bearer a-bearer-token",
					},
				},
			},
			nil,
		},
		{
			"Successfully unmarshalling a HeaderList from an empty slice of bytes",
			make([]byte, 0),
			HeaderList{},
			nil,
		},
		{
			"Successfully unmarshalling a HeaderList from an array of objects",
			[]byte("[{\"key\": \"Content-Type\",\"value\": \"application\\/json\",\"description\": \"The content type\"},{\"key\": \"Authorization\",\"value\": \"Bearer a-bearer-token\",\"description\": \"The Bearer token\"}]"),
			HeaderList{
				Headers: []*Header{
					{
						Key:         "Content-Type",
						Value:       "application/json",
						Description: "The content type",
					},
					{
						Key:         "Authorization",
						Value:       "Bearer a-bearer-token",
						Description: "The Bearer token",
					},
				},
			},
			nil,
		},
		{
			"Failed to unmarshal a HeaderList because of an invalid header",
			[]byte("\"Content-Type\n\""),
			HeaderList{},
			errors.New("invalid header, missing key or value: Content-Type"),
		},
		{
			"Failed to unmarshal a HeaderList because of an unsupported type",
			[]byte("not-a-valid-header-list"),
			HeaderList{},
			errors.New("unsupported type for header list"),
		},
	}

	for _, tc := range cases {

		var hl HeaderList
		err := hl.UnmarshalJSON(tc.bytes)

		assert.Equal(t, tc.expectedHeaderList, hl, tc.scenario)
		assert.Equal(t, tc.expectedError, err, tc.scenario)
	}
}
