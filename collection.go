package postman

import (
	"encoding/json"
	"fmt"
	"io"
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
	Info  Info    `json:"info"`
	Items []Items `json:"item"`
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
	}
}

// AddItem appends an item (Item or ItemGroup) to the existing items slice.
func (c *Collection) AddItem(item Items) {
	c.Items = append(c.Items, item)
}

// AddItemGroup creates a new ItemGroup and appends it to the existing items slice.
func (c *Collection) AddItemGroup(name string) (f *ItemGroup) {
	f = &ItemGroup{
		Name: name,
	}

	c.Items = append(c.Items, f)

	return
}

// Write encodes the Collection struct in JSON and writes it into the provided io.Writer.
func (c *Collection) Write(w io.Writer) (err error) {
	file, err := json.MarshalIndent(c, "", "    ")

	if err != nil {
		return
	}

	_, err = w.Write(file)

	return
}

// ParseCollection parses the content of the provided data stream into a Collection object.
func ParseCollection(r io.Reader) (c *Collection, err error) {

	err = json.NewDecoder(r).Decode(&c)

	return
}

// collectionUnmarshal is used only during unmarshalling process.
// It is used as a temporary object in order to be able to deserialize properly Items objects.
type collectionUnmarshal struct {
	Info  Info          `json:"info"`
	Items []interface{} `json:"item"`
}

// UnmarshalJSON deserializes a JSON into a Collection object.
func (c *Collection) UnmarshalJSON(b []byte) (err error) {

	var collection collectionUnmarshal
	err = json.Unmarshal(b, &collection)

	if err != nil {
		return
	}

	c.Info = collection.Info
	c.Items, err = createItemCollection(collection.Items)

	return
}
