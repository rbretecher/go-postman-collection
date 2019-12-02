package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	cases := []struct {
		method          method
		url             string
		expectedRequest *Request
	}{
		{
			Get,
			"an-url",
			&Request{
				Method: Get,
				URL: &URL{
					Raw: "an-url",
				},
			},
		},
	}

	for _, tc := range cases {
		req := NewRequest(tc.url, tc.method)

		assert.Equal(t, tc.expectedRequest, req)
	}
}
