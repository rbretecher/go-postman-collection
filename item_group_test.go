package postman

import "testing"

import "github.com/stretchr/testify/assert"

func TestAddItemIntoItemGroup(t *testing.T) {
	ig := new(ItemGroup)

	assert.Equal(t, len(ig.Items), 0)

	ig.AddItem(&Item{Name: "Item1"})
	ig.AddItem(&ItemGroup{Name: "Item2"})
	ig.AddItem(&Item{Name: "Item3"})

	assert.Equal(t, ig.Items, []Items{
		&Item{Name: "Item1"},
		&ItemGroup{Name: "Item2"},
		&Item{Name: "Item3"},
	})
}

func TestAddItemGroupIntoItemGroup(t *testing.T) {
	ig := new(ItemGroup)

	assert.Equal(t, len(ig.Items), 0)

	ig.AddItemGroup("new-item-group")
	ig.AddItemGroup("another-new-item-group")

	assert.Equal(t, ig.Items, []Items{
		&ItemGroup{Name: "new-item-group"},
		&ItemGroup{Name: "another-new-item-group"},
	})
}

func TestItemGroupMarsalJSON(t *testing.T) {
	cases := []struct {
		scenario       string
		ig             *ItemGroup
		expectedOutput string
	}{
		{
			"Successfully marshalling an ItemGroup",
			&ItemGroup{
				Name: "an-item-group",
				Items: []Items{
					&Item{
						Name: "a-sub-item",
					},
				},
			},
			"{\"name\":\"an-item-group\",\"item\":[{\"name\":\"a-sub-item\"}]}",
		},
		{
			"Successfully marshalling an ItemGroup without any Items",
			&ItemGroup{
				Name: "an-item-group",
			},
			"{\"name\":\"an-item-group\",\"item\":[]}",
		},
	}

	for _, tc := range cases {
		bytes, _ := tc.ig.MarshalJSON()

		assert.Equal(t, tc.expectedOutput, string(bytes), tc.scenario)
	}
}
