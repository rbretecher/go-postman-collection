package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGroup(t *testing.T) {

	cases := []struct {
		scenario        string
		item            Items
		expectedIsGroup bool
	}{
		{
			"An item with Items is a group",
			Items{
				Name:  "a-name",
				Items: make([]*Items, 0),
			},
			true,
		},
		{
			"An item without Items is not a group",
			Items{
				Name: "a-name",
			},
			false,
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expectedIsGroup, tc.item.IsGroup(), tc.scenario)
	}
}

func TestAddItem(t *testing.T) {
	itemGroup := Items{
		Name:  "A group of items",
		Items: make([]*Items, 0),
	}

	itemGroup.AddItem(&Items{
		Name: "A basic item",
	})

	itemGroup.AddItem(&Items{
		Name:  "A basic group item",
		Items: make([]*Items, 0),
	})

	assert.Equal(
		t,
		Items{
			Name: "A group of items",
			Items: []*Items{
				{
					Name: "A basic item",
				},
				{
					Name:  "A basic group item",
					Items: make([]*Items, 0),
				},
			},
		},
		itemGroup,
	)
}

func TestAddItemGroup(t *testing.T) {
	itemGroup := Items{
		Name:  "A group of items",
		Items: make([]*Items, 0),
	}

	itemGroup.AddItemGroup("an-item-group")
	itemGroup.AddItemGroup("another-item-group")

	assert.Equal(
		t,
		Items{
			Name: "A group of items",
			Items: []*Items{
				{
					Name:  "an-item-group",
					Items: make([]*Items, 0),
				},
				{
					Name:  "another-item-group",
					Items: make([]*Items, 0),
				},
			},
		},
		itemGroup,
	)
}

func TestItemsMarshalJSON(t *testing.T) {
	cases := []struct {
		scenario       string
		item           Items
		expectedOutput string
	}{
		{
			"Successfully marshalling an Item",
			Items{
				ID:   "a-unique-id",
				Name: "an-item",
			},
			"{\"name\":\"an-item\",\"id\":\"a-unique-id\"}",
		},
		{
			"Successfully marshalling a GroupItem",
			Items{
				Name:  "a-group-item",
				Items: make([]*Items, 0),
			},
			"{\"name\":\"a-group-item\",\"item\":[]}",
		},
	}

	for _, tc := range cases {
		bytes, _ := tc.item.MarshalJSON()

		assert.Equal(t, tc.expectedOutput, string(bytes), tc.scenario)
	}
}

func TestCreateItem(t *testing.T) {
	c := CreateItem(Item{
		Name:        "An item",
		Description: "A description",
		Variables: []*Variable{
			{
				Name:  "variable-name",
				Value: "variable-value",
			},
		},
		Event:                   "an-event",
		ProtocolProfileBehavior: "a-protocol-profile-behavior",
		ID:                      "an-id",
		Request: &Request{
			URL: &URL{
				Raw: "http://www.google.fr",
			},
		},
		Response: []*Response{
			{
				Name: "a-response",
			},
		},
	})

	assert.Equal(
		t,
		&Items{
			Name:        "An item",
			Description: "A description",
			Variables: []*Variable{
				{
					Name:  "variable-name",
					Value: "variable-value",
				},
			},
			Event:                   "an-event",
			ProtocolProfileBehavior: "a-protocol-profile-behavior",
			ID:                      "an-id",
			Request: &Request{
				URL: &URL{
					Raw: "http://www.google.fr",
				},
			},
			Response: []*Response{
				{
					Name: "a-response",
				},
			},
		},
		c,
	)
}

func TestCreateItemGroup(t *testing.T) {
	c := CreateItemGroup(ItemGroup{
		Name:        "An item",
		Description: "A description",
		Variables: []*Variable{
			{
				Name:  "variable-name",
				Value: "variable-value",
			},
		},
		Event:                   "an-event",
		ProtocolProfileBehavior: "a-protocol-profile-behavior",
		Items: []*Items{
			{
				Name: "An item",
			},
		},
		Auth: &Auth{
			Type: Basic,
		},
	})

	assert.Equal(
		t,
		&Items{
			Name:        "An item",
			Description: "A description",
			Variables: []*Variable{
				{
					Name:  "variable-name",
					Value: "variable-value",
				},
			},
			Event:                   "an-event",
			ProtocolProfileBehavior: "a-protocol-profile-behavior",
			Items: []*Items{
				{
					Name: "An item",
				},
			},
			Auth: &Auth{
				Type: Basic,
			},
		},
		c,
	)
}
