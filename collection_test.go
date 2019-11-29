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
	suite.Collection = CreateCollection("a-name", "a-desc")
	suite.BasicCollection = &Collection{
		Info: Info{
			Name:        "Go Collection",
			Description: "Awesome description",
			Version:     "v2.1.0",
			Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/",
		},
		Items: []Items{
			&ItemGroup{
				Name: "This is a folder",
				Items: []Items{
					&Item{
						Name: "An item inside a folder",
					},
				},
			},
			&Item{
				Name: "This is a request",
				Request: &Request{
					URL: &URL{
						Raw: "http://www.google.fr",
					},
					Method: Get,
				},
			},
			&Item{
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
			&ItemGroup{
				Name: "An empty folder",
			},
		},
	}
}

func TestCollectionTestSuite(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}

func TestCreateCollection(t *testing.T) {
	c := CreateCollection("a-name", "a-desc")

	assert.Equal(t, &Collection{
		Info: Info{
			Name:        "a-name",
			Description: "a-desc",
			Version:     "v2.1.0",
			Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/",
		},
		Items: []Items{},
	}, c)
}

func (suite *CollectionTestSuite) TestAddItem() {
	suite.Collection.AddItem(new(Item))
	suite.Collection.AddItem(new(ItemGroup))
	suite.Collection.AddItem(new(Item))

	assert.Equal(suite.T(), 3, len(suite.Collection.Items))
}

func (suite *CollectionTestSuite) TestAddItemGroup() {
	suite.Collection.AddItemGroup("new-item-group")
	suite.Collection.AddItemGroup("another-new-item-group")

	if assert.NotNil(suite.T(), suite.Collection.Items) {
		assert.Equal(suite.T(), 2, len(suite.Collection.Items))
		assert.Equal(suite.T(), "new-item-group", suite.Collection.Items[0].(*ItemGroup).Name)
		assert.Equal(suite.T(), "another-new-item-group", suite.Collection.Items[1].(*ItemGroup).Name)
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
			"Basic collection",
			"testdata/basic_collection.json",
			suite.BasicCollection,
			nil,
		},
	}

	for _, tc := range cases {
		file, _ := os.Open(tc.testFile)

		c, err := ParseCollection(file)

		assert.Equal(suite.T(), tc.expectedError, err)
		assert.Equal(suite.T(), tc.expectedCollection, c)
	}
}

func (suite *CollectionTestSuite) TestUnmarshalJSON() {
	cases := []struct {
		scenario           string
		testFile           string
		expectedCollection *Collection
		expectedError      error
	}{
		{
			"Unmarshal valid JSON file should not return any error",
			"testdata/basic_collection.json",
			suite.BasicCollection,
			nil,
		},
		{
			"Unmarshal invalid JSON file should return an error",
			"testdata/malformed_json.json",
			&Collection{},
			assert.AnError,
		},
	}

	for _, tc := range cases {
		c := new(Collection)

		file, err := ioutil.ReadFile(tc.testFile)

		if err != nil {
			suite.Errorf(err, "Could not open test file", tc.scenario)
		}

		err = c.UnmarshalJSON(file)

		if tc.expectedError != nil {
			assert.Error(suite.T(), err, tc.scenario)
		} else {
			assert.NoError(suite.T(), err, tc.scenario)
		}

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
			"Write collection struct into an io.writer",
			suite.BasicCollection,
			"testdata/basic_collection.json",
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

		assert.Equal(suite.T(), string(file), fmt.Sprintf("%s\n", buf.String()))
	}
}
