package postman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const version = "v2.1.0"

// Info stores data about the collection.
type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Schema      string `json:"schema"`
}

// Collection represents a Postman Collection.
type Collection struct {
	Info Info    `json:"info"`
	Item []Items `json:"item"`
}

// CreateCollection returns a new Collection.
func CreateCollection(name string, desc string) *Collection {
	return &Collection{
		Info: Info{
			Name:        name,
			Version:     version,
			Description: desc,
			Schema:      fmt.Sprintf("https://schema.getpostman.com/json/collection/%s/", version),
		},
		Item: make([]Items, 0),
	}
}

func (c *Collection) AddItem(item Items) {
	c.Item = append(c.Item, item)
}

func (c *Collection) AddItemGroup(name string) (f *ItemGroup) {
	f = &ItemGroup{
		Name: name,
		Item: make([]Items, 0),
	}

	c.Item = append(c.Item, f)

	return
}

// Write the collection to a file named by filename.
func (c *Collection) Write(filename string) (err error) {
	file, err := json.MarshalIndent(c, "", "  ")

	if err != nil {
		return
	}

	err = ioutil.WriteFile(filename, file, 0644)

	return
}

// ParseCollection reads the content of the provided file and
//  parse the data into a Collection object.
func ParseCollection(filename string) (c *Collection, err error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(file), &c)

	return
}

// CollectionUnmarshal is used only during unmarshalling process.
// It is used as a temporary object in order to be able to deserialize
//	properly Items objects.
type CollectionUnmarshal struct {
	Info Info          `json:"info"`
	Item []interface{} `json:"item"`
}

// UnmarshalJSON deserializes a JSON into a Collection object.
func (c *Collection) UnmarshalJSON(b []byte) (err error) {

	var collection CollectionUnmarshal
	err = json.Unmarshal(b, &collection)

	if err != nil {
		return
	}

	c.Info = collection.Info
	c.Item, err = createItemCollection(collection.Item)

	return
}
