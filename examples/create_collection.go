package main

import "github.com/rbretecher/go-postman-collection"

func main() {
	c := postman.CreateCollection("Go Collection", "Awesome description")

	c.AddItemGroup("This is a folder").AddItem(&postman.Item{
		Name: "An item inside a folder",
	})

	c.AddItem(&postman.Item{
		Name:    "This is a request",
		Request: postman.NewRequest("http://www.google.fr", "GET"),
	})

	c.AddItemGroup("Empty folder")

	c.Write("postman_collection.json")
}
