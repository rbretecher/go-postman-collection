package postman

import (
	"encoding/json"
	"fmt"
	"io"
)

type version string

const (
	// V210 : v2.1.0
	V210 version = "v2.1.0"
	// V200 : v2.0.0
	V200 version = "v2.0.0"
)

// Info stores data about the collection.
type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Schema      string `json:"schema"`
}

// Collection represents a Postman Collection.
type Collection struct {
	Info      Info        `json:"info"`
	Items     []*Items    `json:"item"`
	Variables []*Variable `json:"variable,omitempty"`
}

// CreateCollection returns a new Collection.
func CreateCollection(name string, desc string) *Collection {
	return &Collection{
		Info: Info{
			Name:        name,
			Description: desc,
		},
	}
}

// AddItem appends an item (Item or ItemGroup) to the existing items slice.
func (c *Collection) AddItem(item *Items) {
	c.Items = append(c.Items, item)
}

// AddItemGroup creates a new ItemGroup and appends it to the existing items slice.
func (c *Collection) AddItemGroup(name string) (f *Items) {
	f = &Items{
		Name:  name,
		Items: make([]*Items, 0),
	}

	c.Items = append(c.Items, f)

	return
}

// Write encodes the Collection struct in JSON and writes it into the provided io.Writer.
func (c *Collection) Write(w io.Writer, v version) (err error) {

	c.Info.Schema = fmt.Sprintf("https://schema.getpostman.com/json/collection/%s/collection.json", string(v))
	setVersionForItems(c.Items, v)

	file, _ := json.MarshalIndent(c, "", "    ")

	_, err = w.Write(file)

	return
}

// Set the version on all structs that have a different behavior depending on the Postman Collection version.
func setVersionForItems(items []*Items, v version) {
	for _, i := range items {
		if i.Auth != nil {
			i.Auth.setVersion(v)
		}
		if i.IsGroup() {
			setVersionForItems(i.Items, v)
		} else {
			if i.Request != nil {
				if i.Request.Auth != nil {
					i.Request.Auth.setVersion(v)
				}
				if i.Request.URL != nil {
					i.Request.URL.setVersion(v)
				}
			}
		}
	}
}

// ParseCollection parses the content of the provided data stream into a Collection object.
func ParseCollection(r io.Reader) (c *Collection, err error) {

	err = json.NewDecoder(r).Decode(&c)

	return
}
