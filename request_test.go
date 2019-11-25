package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	var cases = []struct {
		Method method
		URL    string
	}{
		{Get, "an-url"},
		{Post, "another-url"},
	}

	for _, tc := range cases {
		req := NewRequest(tc.URL, tc.Method)

		assert.Equal(t, tc.Method, req.Method)

		if assert.NotNil(t, req.URL) {
			assert.Equal(t, tc.URL, req.URL.Raw)
		}
	}
}

func TestCreateRequestFromInterfaceWithString(t *testing.T) {
	req, err := createRequestFromInterface("request-from-a-string")

	assert.Nil(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, Get, req.Method)

	if assert.NotNil(t, req.URL) {
		assert.Equal(t, "request-from-a-string", req.URL.Raw)
	}
}

func TestCreateRequestFromInterfaceWithUnsupportedInterface(t *testing.T) {

	var cases = []struct {
		UnsupportedInterface interface{}
	}{
		{666},
		{[]string{"not-a-request"}},
	}

	for _, tc := range cases {
		_, err := createRequestFromInterface(tc.UnsupportedInterface)

		assert.NotNil(t, err)
		assert.Equal(t, "Unsupported interface type", err.Error())
	}
}
