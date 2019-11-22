package postman

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// A ItemGroup is an ordered set of requests.
type ItemGroup struct {
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variable                interface{} `json:"variable,omitempty"`
	Item                    []Items     `json:"item"`
	Event                   interface{} `json:"event,omitempty"`
	Auth                    interface{} `json:"auth,omitempty"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior,omitempty"`
}

func (ig *ItemGroup) IsGroup() bool {
	return true
}

func (ig *ItemGroup) AddItem(item Items) {
	ig.Item = append(ig.Item, item)
}

func (c *ItemGroup) AddItemGroup(name string) (f *ItemGroup) {
	f = &ItemGroup{
		Name: name,
		Item: make([]Items, 0),
	}

	c.Item = append(c.Item, f)

	return
}

func createItemGroupFromMap(m map[string]interface{}) (ig *ItemGroup, err error) {

	config := &mapstructure.DecoderConfig{
		Result: &ig,
		DecodeHook: func(from reflect.Type, to reflect.Type, v interface{}) (interface{}, error) {
			// We need to manually take care of Items as it can be an Item or an ItemGroup.
			if to.Name() == "Items" {
				i, err := createItem(v)
				return i, err
			}

			return v, nil
		},
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(m)

	return
}
