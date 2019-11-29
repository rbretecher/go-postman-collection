# go-postman-collection

[![Test](https://github.com/rbretecher/go-postman-collection/workflows/test/badge.svg)](https://github.com/rbretecher/go-postman-collection/actions?query=workflow=test)
[![Go Report Card](https://goreportcard.com/badge/github.com/rbretecher/go-postman-collection)](https://goreportcard.com/report/github.com/rbretecher/go-postman-collection)
[![GoDoc](https://godoc.org/github.com/rbretecher/go-postman-collection?status.svg)](https://godoc.org/github.com/rbretecher/go-postman-collection)

Go module to work with Postman Collections.

This module aims to provide a simple way to work with Postman collections. Using this module, you can create collections, update them and export them into a the Postman Collection format v2.1.0

### Examples

#### Read a Postman collection

```go
package main

import (
	"os"

	postman "github.com/rbretecher/go-postman-collection"
)

func main() {
	file, err := os.Open("postman_collection.json")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	c, err := postman.ParseCollection(file)

	_ = c
}
```

#### Create and save a Postman collection

```go
package main

import (
	"os"

	postman "github.com/rbretecher/go-postman-collection"
)

func main() {
    c := postman.CreateCollection("My collection", "My awesome collection")

	c.AddItemGroup("A folder").AddItem(&postman.Item{
		Name:    "This is a request",
		Request: postman.NewRequest("http://www.google.fr", postman.Get),
    })

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
```

### Current support

Development is under progress and for now it only supports partially Postman Collection format v2.1.0

|  Object            | Supported |
| ------------------ | --------- |
| Collection         | Yes       |
| ItemGroup (Folder) | Yes       |
| Item               | Yes       |
| Request            | Yes       |
| Response           | No        |
| Event              | No        |
| Variable           | No        |
| Auth               | Yes       |
