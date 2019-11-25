package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	var tests = []struct {
		Method               method
		URL                  string
		ExpectedMethodString string
	}{
		{Get, "an-url", "GET"},
		{Post, "another-url", "POST"},
	}

	for _, test := range tests {
		req := NewRequest(test.URL, test.Method)

		assert.Equal(t, test.ExpectedMethodString, req.Method)

		if assert.NotNil(t, req.URL) {
			assert.Equal(t, test.URL, req.URL.Raw)
		}
	}
}

func TestCreateRequestFromInterfaceWithString(t *testing.T) {
	req, err := createRequestFromInterface("request-from-a-string")

	assert.Nil(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, "GET", req.Method)

	if assert.NotNil(t, req.URL) {
		assert.Equal(t, "request-from-a-string", req.URL.Raw)
	}
}

func TestCreateRequestFromInterfaceWithUnsupportedInterface(t *testing.T) {

	var tests = []struct {
		UnsupportedInterface interface{}
	}{
		{666},
		{[]string{"not-a-request"}},
	}

	for _, test := range tests {
		_, err := createRequestFromInterface(test.UnsupportedInterface)

		assert.NotNil(t, err)
		assert.Equal(t, "Unsupported interface type", err.Error())
	}
}
