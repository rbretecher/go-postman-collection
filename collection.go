package main

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

func (c *Collection) AddFolder(name string) (f *Folder) {
	f = &Folder{
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
