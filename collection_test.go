package postman

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CollectionTestSuite struct {
	suite.Suite
	Collection     *Collection
	V200Collection *Collection
	V210Collection *Collection
}

func (suite *CollectionTestSuite) SetupTest() {
	suite.Collection = CreateCollection("Postman collection", "v2.1.0")
	suite.V200Collection = &Collection{
		Info: Info{
			Name:        "Go Collection",
			Description: "Awesome description",
			Version:     "v2.0.0",
			Schema:      "https://schema.getpostman.com/json/collection/v2.0.0/collection.json",
		},
		Items: []*Items{
			{
				Name: "This is a folder",
				Items: []*Items{
					{
						Name: "An item inside a folder",
					},
				},
				Variables: []*Variable{
					{
						Name:  "api-key",
						Value: "abcd1234",
					},
				},
			},
			{
				Name: "This is a request",
				Request: &Request{
					URL: &URL{
						Raw: "http://www.google.fr",
					},
					Method: Get,
				},
			},
			{
				Name: "JSON-RPC Request",
				Request: &Request{
					URL: &URL{
						Raw: "https://gurujsonrpc.appspot.com/guru",
					},
					Auth: &Auth{
						Type: Basic,
						Basic: []*AuthParam{
							{
								Key:   "password",
								Value: "my-password",
							},
						},
					},
					Method: Post,
					Header: []*Header{
						{
							Key:   "Content-Type",
							Value: "application/json",
						},
					},
					Body: &Body{
						Mode:    "raw",
						Raw:     "{\"aKey\":\"a-value\"}",
						Options: &BodyOptions{BodyOptionsRaw{Language: "json"}},
					},
				},
			},
			{
				Name:  "An empty folder",
				Items: make([]*Items, 0),
			},
		},
		Variables: []*Variable{
			{
				Name:  "a-global-collection-variable",
				Value: "a-global-value",
			},
		},
	}
	suite.V210Collection = &Collection{
		Info: Info{
			Name:        "Go Collection",
			Description: "Awesome description",
			Version:     "v2.1.0",
			Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		Items: []*Items{
			{
				Name: "This is a folder",
				Items: []*Items{
					{
						Name: "An item inside a folder",
					},
				},
				Variables: []*Variable{
					{
						Name:  "api-key",
						Value: "abcd1234",
					},
				},
				Auth: &Auth{
					Type: Bearer,
					Bearer: []*AuthParam{
						{
							Key:   "token",
							Value: "a-bearer-token",
							Type:  "string",
						},
					},
				},
			},
			{
				Name: "This is a request",
				Request: &Request{
					URL: &URL{
						Raw: "http://www.google.fr",
					},
					Method: Get,
				},
			},
			{
				Name: "JSON-RPC Request",
				Request: &Request{
					URL: &URL{
						Raw: "https://gurujsonrpc.appspot.com/guru",
						Variables: []*Variable{
							{
								Name:  "an-url-variable",
								Value: "an-url-variable-value",
							},
						},
					},
					Auth: &Auth{
						Type: Basic,
						Basic: []*AuthParam{
							{
								Key:   "password",
								Value: "my-password",
								Type:  "string",
							},
						},
					},
					Method: Post,
					Header: []*Header{
						{
							Key:   "Content-Type",
							Value: "application/json",
						},
					},
					Body: &Body{
						Mode:    "raw",
						Raw:     "{\"aKey\":\"a-value\"}",
						Options: &BodyOptions{BodyOptionsRaw{Language: "json"}},
					},
				},
			},
			{
				Name:  "An empty folder",
				Items: make([]*Items, 0),
			},
		},
		Variables: []*Variable{
			{
				Name:  "a-global-collection-variable",
				Value: "a-global-value",
			},
		},
	}
}

func TestCollectionTestSuite(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}

func TestCreateCollection(t *testing.T) {
	c := CreateCollection("a-name", "a-desc")

	assert.Equal(
		t,
		&Collection{
			Info: Info{
				Name:        "a-name",
				Description: "a-desc",
			},
		},
		c,
	)
}

func (suite *CollectionTestSuite) TestAddItemIntoCollection() {
	suite.Collection.AddItem(&Items{Name: "Item1"})
	suite.Collection.AddItem(&Items{Name: "Item2"})
	suite.Collection.AddItem(&Items{Name: "Item3"})

	assert.Equal(
		suite.T(),
		[]*Items{
			{Name: "Item1"},
			{Name: "Item2"},
			{Name: "Item3"},
		},
		suite.Collection.Items,
	)
}

func (suite *CollectionTestSuite) TestAddItemGroupIntoCollection() {
	suite.Collection.AddItemGroup("new-item-group")
	suite.Collection.AddItemGroup("another-new-item-group")

	if assert.NotNil(suite.T(), suite.Collection.Items) {

		assert.Equal(
			suite.T(),
			[]*Items{
				{Name: "new-item-group", Items: make([]*Items, 0)},
				{Name: "another-new-item-group", Items: make([]*Items, 0)},
			},
			suite.Collection.Items,
		)
	}
}

func (suite *CollectionTestSuite) TestParseCollection() {
	cases := []struct {
		scenario           string
		testFile           string
		expectedCollection *Collection
		expectedError      error
	}{
		{
			"v2.0.0 collection",
			"testdata/collection_v2.0.0.json",
			suite.V200Collection,
			nil,
		},
		{
			"v2.1.0 collection",
			"testdata/collection_v2.1.0.json",
			suite.V210Collection,
			nil,
		},
	}

	for _, tc := range cases {
		file, _ := os.Open(tc.testFile)

		c, err := ParseCollection(file)
		assert.Equal(suite.T(), tc.expectedError, err, tc.scenario)
		assert.Equal(suite.T(), tc.expectedCollection, c, tc.scenario)
	}
}

func (suite *CollectionTestSuite) TestWriteCollection() {
	cases := []struct {
		scenario       string
		testCollection *Collection
		expectedFile   string
		expectedError  error
	}{
		{
			"v2.1.0 collection",
			suite.V210Collection,
			"testdata/collection_v2.1.0.json",
			nil,
		},
	}

	for _, tc := range cases {
		var buf bytes.Buffer

		tc.testCollection.Write(&buf, V210)

		file, err := ioutil.ReadFile(tc.expectedFile)

		if err != nil {
			suite.Errorf(err, "Could not open test file")
		}

		assert.Equal(suite.T(), string(file), fmt.Sprintf("%s\n", buf.String()), tc.scenario)
	}
}

func (suite *CollectionTestSuite) TestSimplePOSTItem() {
	c := CreateCollection("Test Collection", "My Test Collection")

	file, err := os.Create("postman_collection.json")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), file)

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
		Options: &BodyOptions{BodyOptionsRaw{Language: "json"}},
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
	assert.Nil(suite.T(), err)

	err = os.Remove("postman_collection.json")
	assert.Nil(suite.T(), err)
}

func (suite *CollectionTestSuite) TestSimpleGETItem() {
	c := CreateCollection("Test Collection", "My Test Collection")

	file, err := os.Create("postman_collection.json")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), file)

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
	assert.Nil(suite.T(), err)

	err = os.Remove("postman_collection.json")
	assert.Nil(suite.T(), err)
}
