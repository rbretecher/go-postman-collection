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
	Items                   []Items     `json:"item"`
	Event                   interface{} `json:"event,omitempty"`
	Auth                    *Auth       `json:"auth,omitempty"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior,omitempty"`
}

func (ig *ItemGroup) IsGroup() bool {
	return true
}

func (ig *ItemGroup) AddItem(item Items) {
	ig.Items = append(ig.Items, item)
}

func (c *ItemGroup) AddItemGroup(name string) (f *ItemGroup) {
	f = &ItemGroup{
		Name:  name,
		Items: make([]Items, 0),
	}

	c.Items = append(c.Items, f)

	return
}

func createItemGroupFromMap(m map[string]interface{}) (ig *ItemGroup, err error) {
	config := &mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &ig,
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
		return
	}

	err = decoder.Decode(m)

	return
}
