package postman

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequest(t *testing.T) {
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
		req := &Request{
			URL: &URL{
				Raw: tc.url,
			},
			Method: tc.method,
		}

		assert.Equal(t, tc.expectedRequest, req)
	}
}

func TestRequestMarshalJSON(t *testing.T) {
	cases := []struct {
		scenario       string
		req            Request
		expectedOutput string
	}{
		{
			"Successfully marshalling a Request as a string",
			Request{
				Method: Get,
				URL: &URL{
					Raw:     "http://www.google.fr",
					version: V200,
				},
			},
			"\"http://www.google.fr\"",
		},
		{
			"Successfully marshalling a Request as an object (v2.0.0)",
			Request{
				Method: Post,
				URL: &URL{
					Raw:     "http://www.google.fr",
					version: V200,
				},
				Body: &Body{
					Mode:    "raw",
					Raw:     "raw-content",
					Options: BodyOptions{BodyOptionsRaw{Language: "json"}},
				},
			},
			"{\"url\":\"http://www.google.fr\",\"method\":\"POST\",\"body\":{\"mode\":\"raw\",\"raw\":\"raw-content\",\"options\":{\"raw\":{\"language\":\"json\"}}}}",
		},
		{
			"Successfully marshalling a Request as an object (v2.1.0)",
			Request{
				Method: Post,
				URL: &URL{
					Raw:     "http://www.google.fr",
					version: V210,
				},
				Body: &Body{
					Mode: "raw",
					Raw:  "raw-content",
				},
			},
			"{\"url\":{\"raw\":\"http://www.google.fr\"},\"method\":\"POST\",\"body\":{\"mode\":\"raw\",\"raw\":\"raw-content\"qq}}",
		},
	}

	for _, tc := range cases {
		bytes, _ := tc.req.MarshalJSON()

		assert.Equal(t, tc.expectedOutput, string(bytes), tc.scenario)
	}
}

func TestRequestUnmarshalJSON(t *testing.T) {
	cases := []struct {
		scenario        string
		bytes           []byte
		expectedRequest Request
		expectedError   error
	}{
		{
			"Successfully unmarshalling a Request as a string",
			[]byte("\"http://www.google.fr\""),
			Request{
				Method: Get,
				URL: &URL{
					Raw: "http://www.google.fr",
				},
			},
			nil,
		},
		{
			"Successfully unmarshalling a Request URL with some content",
			[]byte("{\"url\":\"http://www.google.fr\",\"body\":{\"mode\":\"raw\",\"raw\":\"awesome-body\"}}"),
			Request{
				URL: &URL{
					Raw: "http://www.google.fr",
				},
				Body: &Body{
					Mode: "raw",
					Raw:  "awesome-body",
				},
			},
			nil,
		},
		{
			"Failed to unmarshal a Request because of an unsupported type",
			[]byte("not-a-valid-request"),
			Request{},
			errors.New("Unsupported type"),
		},
	}

	for _, tc := range cases {

		var r Request
		err := r.UnmarshalJSON(tc.bytes)

		assert.Equal(t, tc.expectedRequest, r, tc.scenario)
		assert.Equal(t, tc.expectedError, err, tc.scenario)
	}
}

func TestSimplePOSTItem(t *testing.T) {
	c := CreateCollection("Test Collection", "My Test Collection")

	file, err := os.Create("postman_collection.json")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	defer file.Close()

	pURL := URL{
		Raw:      "https://test.com",
		Protocol: "https",
		Host:     []string{"test", "com"},
	}

	headers := []*Header{{
		Key:   "h1",
		Value: "h1-value",
	}}

	pBody := Body{
		Mode:    "raw",
		Raw:     "{\"a\":\"1234\",\"b\":123}",
		Options: BodyOptions{BodyOptionsRaw{Language: "json"}},
	}

	pReq := Request{
		Method: Post,
		URL:    &pURL,
		Header: headers,
		Body:   &pBody,
	}

	cr := Request{
		Method: Post,
		URL:    &pURL,
		Header: pReq.Header,
		Body:   pReq.Body,
	}

	item := CreateItem(Item{
		Name:    "Test-POST",
		Request: &cr,
	})

	c.AddItemGroup("grp1").AddItem(item)

	err = c.Write(file, V210)
	assert.Nil(t, err)

	err = os.Remove("postman_collection.json")
	assert.Nil(t, err)
}

func TestSimpleGETItem(t *testing.T) {
	c := CreateCollection("Test Collection", "My Test Collection")

	file, err := os.Create("postman_collection.json")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	defer file.Close()

	m1 := map[string]interface{}{"key": "param1", "value": "value1"}
	m2 := map[string]interface{}{"key": "param2", "value": "value2"}

	var arrMaps []map[string]interface{}
	arrMaps = append(arrMaps, m1)
	arrMaps = append(arrMaps, m2)

	pURL := URL{
		Raw:      "https://test.com?a=3",
		Protocol: "https",
		Host:     []string{"test", "com"},
		Query:    arrMaps,
	}

	headers := []*Header{}
	headers = append(headers, &Header{
		Key:   "h1",
		Value: "h1-value",
	})
	headers = append(headers, &Header{
		Key:   "h2",
		Value: "h2-value",
	})

	pReq := Request{
		Method: Get,
		URL:    &pURL,
		Header: headers,
	}

	item := CreateItem(Item{
		Name:    "Test-GET",
		Request: &pReq,
	})

	c.AddItemGroup("grp1").AddItem(item)

	err = c.Write(file, V210)
	assert.Nil(t, err)

	err = os.Remove("postman_collection.json")
	assert.Nil(t, err)
}
