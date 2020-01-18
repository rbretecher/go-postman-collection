# go-postman-collection

[![Build Status](https://travis-ci.org/rbretecher/go-postman-collection.svg?branch=master)](https://travis-ci.org/rbretecher/go-postman-collection)
[![Report](https://goreportcard.com/badge/github.com/rbretecher/go-postman-collection)](https://goreportcard.com/report/github.com/rbretecher/go-postman-collection)
[![Code coverage](https://codecov.io/gh/rbretecher/go-postman-collection/branch/master/graph/badge.svg)](https://codecov.io/gh/rbretecher/go-postman-collection)

Go module to work with Postman Collections.

This module aims to provide a simple way to work with Postman collections. Using this module, you can create collections, update them and export them into the Postman Collection format v2 (compatible with Insomnia)

Postman Collections are a group of saved requests you can organize into folders. For more information about Postman Collections, you can visit the [official documentation](https://www.getpostman.com/collection).

## Examples

### Collections

#### Read a Postman Collection

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

#### Create and save a Postman Collection

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

### Items

`Items` are the basic unit for a Postman collection, it can either be a request (`Item`) or a folder (`ItemGroup`).

```go
// Create a simple item.
item := postman.CreateItem(postman.Item{
    Name:    "A basic request",
    Request: postman.NewRequest("http://www.google.fr", postman.Get),
})

// Create a simple folder.
folder := postman.CreateItemGroup(postman.ItemGroup{
    Name: "A folder",
})

// Add the item to the folder
folder.AddItem(item)
```

### Requests

```go
// Basic request
req := postman.NewRequest("http://www.google.fr", postman.Get)

// Complex request
req := postman.Request{
    URL: &postman.URL{
        Raw: "http://www.google.fr",
    },
    Method: postman.Post,
    Body: &postman.Body{
        Mode: "raw",
        Raw:  "{\"key\": \"value\"}",
    },
}
```

### Auth

`Auth` can be added to a `Request` or an `ItemGroup`.

```go
// Create basic auth with username and password
auth := postman.CreateAuth(postman.Basic, postman.CreateAuthParam("username", "password"))
```

## Current support

For now, it does not offer support for `Response` and `Event` objects. Feel free to submit a pull request if you want to add support for one of those objects.

|  Object            | v2.0.0 | v2.1.0 |
| ------------------ | ------ | ------ |
| Collection         | Yes    | Yes    |
| ItemGroup (Folder) | Yes    | Yes    |
| Item               | Yes    | Yes    |
| Request            | Yes    | Yes    |
| Response           | No     | No     |
| Event              | No     | No     |
| Variable           | Yes    | Yes    |
| Auth               | Yes    | Yes    |
