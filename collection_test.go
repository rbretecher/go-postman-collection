package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CollectionTestSuite struct {
	suite.Suite
	Collection *Collection
}

func (suite *CollectionTestSuite) SetupTest() {
	suite.Collection = CreateCollection("a-name", "a-desc")
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
