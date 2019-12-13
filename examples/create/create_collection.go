package main

import (
	"os"

	postman "github.com/rbretecher/go-postman-collection"
)

func main() {
	c := postman.CreateCollection("Go Collection", "Awesome description", postman.V210)

	c.AddItemGroup("This is a folder").AddItem(&postman.Items{
		Name: "An item inside a folder",
	})

	c.AddItem(&postman.Items{
		Name:    "This is a request",
		Request: postman.NewRequest("http://www.google.fr", postman.Get),
	})

	r := &postman.Request{
		URL: &postman.URL{
			Raw: "https://gurujsonrpc.appspot.com/guru",
		},
		Method: "POST",
		Header: []*postman.Header{
			{
				Key:   "Content-Type",
				Value: "application/json",
			},
		},
		Body: &postman.Body{
			Mode: "raw",
			Raw:  "{\"aKey\":\"a-value\"}",
		},
	}

	c.AddItem(&postman.Items{
		Name:    "JSON-RPC Request",
		Request: r,
	})

	c.AddItemGroup("Empty folder")

	file, err := os.Create("postman_collection.json")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	err = c.Write(file)

	if err != nil {
		panic(err)
	}
}
