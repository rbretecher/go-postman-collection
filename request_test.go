package postman

import (
	"errors"
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

func TestCreateRequestFromInterface(t *testing.T) {
	cases := []struct {
		scenario        string
		i               interface{}
		expectedRequest *Request
		expectedError   error
	}{
		{
			"Succesfully creating a request from a string",
			"http://www.google.fr",
			&Request{
				Method: Get,
				URL: &URL{
					Raw: "http://www.google.fr",
				},
			},
			nil,
		},
		{
			"Succesfully creating a request from a map[string]interface{}",
			map[string]interface{}{
				"method": "POST",
				"header": []map[string]interface{}{
					{
						"key":   "Content-Type",
						"value": "application/json",
					},
				},
				"url": map[string]interface{}{
					"raw":      "https://gurujsonrpc.appspot.com/guru",
					"protocol": "https",
					"host":     []string{"gurujsonrpc", "appspot", "com"},
					"path":     []string{"path"},
				},
				"body": map[string]interface{}{
					"mode": "raw",
					"raw":  "some-raw-body",
				},
			},
			&Request{
				Method: Post,
				Header: []*Header{
					{
						Key:   "Content-Type",
						Value: "application/json",
					},
				},
				URL: &URL{
					Raw:      "https://gurujsonrpc.appspot.com/guru",
					Protocol: "https",
					Host:     []string{"gurujsonrpc", "appspot", "com"},
					Path:     []string{"path"},
				},
				Body: &Body{
					Mode: "raw",
					Raw:  "some-raw-body",
				},
			},
			nil,
		},
		{
			"Request from an unsupported interface",
			[]string{"not-a-request"},
			nil,
			errors.New("Unsupported interface type"),
		},
	}

	for _, tc := range cases {
		req, err := createRequestFromInterface(tc.i)

		assert.Equal(t, tc.expectedError, err, tc.scenario)
		assert.Equal(t, tc.expectedRequest, req, tc.scenario)
	}
}
