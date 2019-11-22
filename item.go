package postman

import "github.com/mitchellh/mapstructure"

// An Item is an entity which contain an actual HTTP request, and sample responses attached to it.
type Item struct {
	ID                      string      `json:"id"`
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variable                interface{} `json:"variable,omitempty"`
	Event                   interface{} `json:"event,omitempty"`
	Request                 *Request    `json:"request,omitempty"`
	Response                interface{} `json:"response,omitempty"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior,omitempty"`
}

func (i *Item) IsGroup() bool {
	return false
}

func createItemFromMap(m map[string]interface{}) (item *Item, err error) {
	err = mapstructure.Decode(m, &item)

	return
}
