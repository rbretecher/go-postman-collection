package postman

import (
	"encoding/json"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// A ItemGroup is an ordered set of requests.
type ItemGroup struct {
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variable                interface{} `json:"variable,omitempty"`
	Items                   []Items     `json:"item"`
	Event                   interface{} `json:"event,omitempty"`
	Auth                    *Auth       `json:"auth,omitempty"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior,omitempty"`
}

type marshalledItemGroup ItemGroup

// IsGroup returns true as an ItemGroup is a group.
func (ig ItemGroup) IsGroup() bool {
	return true
}

// AddItem appends an item (Item or ItemGroup) to the existing items slice.
func (ig *ItemGroup) AddItem(item Items) {
	ig.Items = append(ig.Items, item)
}

// AddItemGroup creates a new ItemGroup and appends it to the existing items slice.
func (ig *ItemGroup) AddItemGroup(name string) (f *ItemGroup) {
	f = &ItemGroup{
		Name:  name,
		Items: make([]Items, 0),
	}

	ig.Items = append(ig.Items, f)

	return
}

func decodeItemGroup(m map[string]interface{}) (ig *ItemGroup, err error) {
	config := &mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &ig,
		DecodeHook: func(from reflect.Type, to reflect.Type, v interface{}) (interface{}, error) {
			// We need to manually take care of Items as it can be an Item or an ItemGroup.
			if to.Name() == "Items" {
				i, err := createItemFromInterface(v)
				return i, err
			}

			return v, nil
		},
	}

	decoder, err := mapstructure.NewDecoder(config)

	if err != nil {
		return
	}

	err = decoder.Decode(m)

	return
}

// MarshalJSON returns JSON encoding of the ItemGroup.
// If the Items slice is nil, it creates an empty one to avoid `null` value in the JSON.
func (ig *ItemGroup) MarshalJSON() ([]byte, error) {

	items := ig.Items
	if ig.Items == nil {
		items = make([]Items, 0)
	}

	return json.Marshal(marshalledItemGroup{
		Name:                    ig.Name,
		Description:             ig.Description,
		Variable:                ig.Variable,
		Items:                   items,
		Event:                   ig.Event,
		Auth:                    ig.Auth,
		ProtocolProfileBehavior: ig.ProtocolProfileBehavior,
	})
}
