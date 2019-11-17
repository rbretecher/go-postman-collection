package main

// Items are the basic unit for a Postman collection.
// It can either be a request or a folder.
type Items interface {
	getName() string
}

// An Item is an entity which contain an actual HTTP request, and sample responses attached to it.
type Item struct {
	ID                      string      `json:"id"`
	Name                    string      `json:"name"`
	Description             string      `json:"description"`
	Variable                interface{} `json:"variable"`
	Event                   interface{} `json:"event"`
	Request                 interface{} `json:"request"`
	Response                interface{} `json:"response"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior"`
}

func (i *Item) getName() string {
	return i.Name
}

// A Folder is an ordered set of requests.
type Folder struct {
	Name                    string      `json:"name"`
	Description             string      `json:"description"`
	Variable                interface{} `json:"variable"`
	Item                    []Items     `json:"item"`
	Event                   interface{} `json:"event"`
	Auth                    interface{} `json:"auth"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior"`
}

func (f *Folder) getName() string {
	return f.Name
}

func (f *Folder) AddItem(item Items) {
	f.Item = append(f.Item, item)
}

func (c *Folder) AddFolder(name string) (f *Folder) {
	f = &Folder{
		Name: name,
		Item: make([]Items, 0),
	}

	c.Item = append(c.Item, f)

	return
}
