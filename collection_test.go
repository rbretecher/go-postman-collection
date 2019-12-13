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
	Collection      *Collection
	BasicCollection *Collection
}

func (suite *CollectionTestSuite) SetupTest() {
	suite.Collection = CreateCollection("Postman collection", "v2.1.0", V210)
	suite.BasicCollection = &Collection{
		Info: Info{
			Name:        "Go Collection",
			Description: "Awesome description",
			Version:     "v2.1.0",
			Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		Items: []*Items{
			&Items{
				Name: "This is a folder",
				Items: []*Items{
					&Items{
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
			&Items{
				Name: "This is a request",
				Request: &Request{
					URL: &URL{
						Raw: "http://www.google.fr",
					},
					Method: Get,
				},
			},
			&Items{
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
							{
								Key:   "username",
								Value: "my-username",
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
						Mode: "raw",
						Raw:  "{\"aKey\":\"a-value\"}",
					},
				},
			},
			&Items{
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
	c := CreateCollection("a-name", "a-desc", V210)

	assert.Equal(
		t,
		&Collection{
			version: V210,
			Info: Info{
				Name:        "a-name",
				Description: "a-desc",
				Version:     "v2.1.0",
				Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
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
			&Items{Name: "Item1"},
			&Items{Name: "Item2"},
			&Items{Name: "Item3"},
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
				&Items{Name: "new-item-group", Items: make([]*Items, 0)},
				&Items{Name: "another-new-item-group", Items: make([]*Items, 0)},
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
			"v2.1.0 collection",
			"testdata/collection_v2.1.0.json",
			suite.BasicCollection,
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
			suite.BasicCollection,
			"testdata/collection_v2.1.0.json",
			nil,
		},
	}

	for _, tc := range cases {
		var buf bytes.Buffer

		tc.testCollection.Write(&buf)

		file, err := ioutil.ReadFile(tc.expectedFile)

		if err != nil {
			suite.Errorf(err, "Could not open test file")
		}

		assert.Equal(suite.T(), string(file), fmt.Sprintf("%s\n", buf.String()), tc.scenario)
	}
}
