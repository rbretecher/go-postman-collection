package postman

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGroup(t *testing.T) {
	cases := []struct {
		scenario       string
		i              Items
		expectedOutput bool
	}{
		{
			"Item is not a group",
			new(Item),
			false,
		},
		{
			"ItemGroup is a group",
			new(ItemGroup),
			true,
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expectedOutput, tc.i.IsGroup())
	}
}

func TestCreateItemCollection(t *testing.T) {

	cases := []struct {
		scenario               string
		i                      []interface{}
		expectedItemCollection []Items
		expectedError          error
	}{
		{
			"Successfully creating item collection from compatible interface",
			[]interface{}{
				map[string]interface{}{
					"name": "An Item",
				},
				map[string]interface{}{
					"name": "An ItemGroup",
					"item": nil,
				},
			},
			[]Items{
				&Item{
					Name: "An Item",
				},
				&ItemGroup{
					Name: "An ItemGroup",
				},
			},
			nil,
		},
		{
			"Failed to create item collection because of an incompatible interface",
			[]interface{}{
				"not-a-valid-item",
				"not-a-valid-item-group",
			},
			nil,
			errors.New("Unsupported interface"),
		},
	}

	for _, tc := range cases {
		items, err := createItemCollection(tc.i)

		assert.Equal(t, tc.expectedError, err, tc.scenario)
		assert.Equal(t, tc.expectedItemCollection, items, tc.scenario)
	}
}
