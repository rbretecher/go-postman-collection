package postman

// Items are the basic unit for a Postman collection.
// It can either be a request (Item) or a folder (ItemGroup).
type Items interface {
	getName() string
}

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

func (i *Item) getName() string {
	return i.Name
}

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

func (f *ItemGroup) getName() string {
	return f.Name
}

func (f *ItemGroup) AddItem(item Items) {
	f.Item = append(f.Item, item)
}

func (c *ItemGroup) AddItemGroup(name string) (f *ItemGroup) {
	f = &ItemGroup{
		Name: name,
		Item: make([]Items, 0),
	}

	c.Item = append(c.Item, f)

	return
}
