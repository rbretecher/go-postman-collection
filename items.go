package main

// Items are the basic unit for a Postman collection.
// It can either be a request or a folder.
type Items struct {
	Name                    string      `json:"name"`
	Description             string      `json:"description"`
	Variable                interface{} `json:"variable"`
	Event                   interface{} `json:"event"`
	ProtocolProfileBehavior interface{} `json:"protocolProfileBehavior"`
}

// An Item is an entity which contain an actual HTTP request, and sample responses attached to it.
type Item struct {
	Items
	ID       string      `json:"id"`
	Request  interface{} `json:"request"`
	Response interface{} `json:"response"`
}

// A Folder is an ordered set of requests.
type Folder struct {
	Items
	Item []Items     `json:"item"`
	Auth interface{} `json:"auth"`
}
