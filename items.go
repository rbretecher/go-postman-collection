package postman

import "encoding/json"

// Items are the basic unit for a Postman collection.
// It can either be a request (Item) or a folder (ItemGroup).
type Items struct {
	// Common fields.
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variables               []*Variable `json:"variable,omitempty"`
	Event                   []*Event    `json:"event,omitempty"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior,omitempty"`
	// Fields specific to Item
	ID        string      `json:"id,omitempty"`
	Request   *Request    `json:"request,omitempty"`
	Responses []*Response `json:"response,omitempty"`
	// Fields specific to ItemGroup
	Items []*Items `json:"item"`
	Auth  *Auth    `json:"auth,omitempty"`
}

// An Item is an entity which contain an actual HTTP request, and sample responses attached to it.
type Item struct {
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variables               []*Variable `json:"variable,omitempty"`
	Event                   []*Event    `json:"event,omitempty"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior,omitempty"`
	ID                      string      `json:"id,omitempty"`
	Request                 *Request    `json:"request,omitempty"`
	Responses               []*Response `json:"response,omitempty"`
}

// A ItemGroup is an ordered set of requests.
type ItemGroup struct {
	Name                    string      `json:"name"`
	Description             string      `json:"description,omitempty"`
	Variables               []*Variable `json:"variable,omitempty"`
	Event                   []*Event    `json:"event,omitempty"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior,omitempty"`
	Items                   []*Items    `json:"item"`
	Auth                    *Auth       `json:"auth,omitempty"`
}

// CreateItem is a helper to create a new Item.
func CreateItem(i Item) *Items {
	return &Items{
		Name:                    i.Name,
		Description:             i.Description,
		Variables:               i.Variables,
		Event:                   i.Event,
		ProtocolProfileBehavior: i.ProtocolProfileBehavior,
		ID:                      i.ID,
		Request:                 i.Request,
		Responses:               i.Responses,
	}
}

// CreateItemGroup is a helper to create a new ItemGroup.
func CreateItemGroup(ig ItemGroup) *Items {
	return &Items{
		Name:                    ig.Name,
		Description:             ig.Description,
		Variables:               ig.Variables,
		Event:                   ig.Event,
		ProtocolProfileBehavior: ig.ProtocolProfileBehavior,
		Items:                   ig.Items,
		Auth:                    ig.Auth,
	}
}

// IsGroup returns false as an Item is not a group.
func (i Items) IsGroup() bool {
	if i.Items != nil {
		return true
	}

	return false
}

// AddItem appends an item to the existing items slice.
func (i *Items) AddItem(item *Items) {
	i.Items = append(i.Items, item)
}

// AddItemGroup creates a new Item folder and appends it to the existing items slice.
func (i *Items) AddItemGroup(name string) (f *Items) {
	f = &Items{
		Name:  name,
		Items: make([]*Items, 0),
	}

	i.Items = append(i.Items, f)

	return
}

// MarshalJSON returns the JSON encoding of an Item/ItemGroup.
func (i Items) MarshalJSON() ([]byte, error) {

	if i.IsGroup() {
		return json.Marshal(ItemGroup{
			Name:                    i.Name,
			Description:             i.Description,
			Variables:               i.Variables,
			Event:                   i.Event,
			ProtocolProfileBehavior: i.ProtocolProfileBehavior,
			Items:                   i.Items,
			Auth:                    i.Auth,
		})
	}

	return json.Marshal(Item{
		Name:                    i.Name,
		Description:             i.Description,
		Variables:               i.Variables,
		Event:                   i.Event,
		ProtocolProfileBehavior: i.ProtocolProfileBehavior,
		ID:                      i.ID,
		Request:                 i.Request,
		Responses:               i.Responses,
	})
}
