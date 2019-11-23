package main

import (
	postman "github.com/rbretecher/go-postman-collection"
)

func main() {
	c := postman.CreateCollection("Go Collection", "Awesome description")

	c.AddItemGroup("This is a folder").AddItem(&postman.Item{
		Name: "An item inside a folder",
	})

	c.AddItem(&postman.Item{
		Name:    "This is a request",
		Request: postman.NewRequest("http://www.google.fr", postman.GET),
	})

	r := &postman.Request{
		URL: postman.URL{
			Raw: "https://gurujsonrpc.appspot.com/guru",
		},
		Method: "POST",
		Header: []*postman.Header{
			&postman.Header{
				Key:   "Content-Type",
				Value: "application/json",
			},
		},
		Body: &postman.Body{
			Mode: "raw",
			Raw:  "{\"aKey\":\"a-value\"}",
		},
	}

	c.AddItem(&postman.Item{
		Name:    "JSON-RPC Request",
		Request: r,
	})

	c.AddItemGroup("Empty folder")

	c.Write("postman_collection.json")
}
