# go-postman-collection

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/rbretecher/go-postman-collection)
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
        Request: Request{
            URL: &URL{
                Raw: "http://www.google.fr",
            },
            Method: postman.Get,
        },
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
    Request: Request{
        URL: &URL{
            Raw: "http://www.google.fr",
        },
        Method: postman.Get,
    }
})

// Create a simple folder.
folder := postman.CreateItemGroup(postman.ItemGroup{
    Name: "A folder",
})

// Add the item to the folder
folder.AddItem(item)
```

### Request

Part of the `Item`, a `Request` represents an HTTP request.

```go
// Basic request
req := Request{
    URL: &URL{
        Raw: "http://www.google.fr",
    },
    Method: postman.Get,
}

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

### Variable

`Variable` can be added to `Collection`, `Item`, `ItemGroup` and `URL`.

```go
v := postman.CreateVariable("env", "prod")
```

### Event

`Event` can be added to `Collection` and run against any `Item` in the entire `Collection`.

```go
scripts := []string{
		`
pm.test("Status code is 200", function () {
	pm.response.to.have.status(200);
});`,
		`
pm.test("Response time is less than 200ms", function () {
    pm.expect(pm.response.responseTime).to.be.below(200);
});`,
}

e := CreateEvent(EventType(EventType_Test), "text/javascript", scripts)
```

## Current support

For now, it does not offer support for `Response` object. Feel free to submit a pull request if you want to add support for one of those objects.

|  Object            | v2.0.0 | v2.1.0 |
| ------------------ | ------ | ------ |
| Collection         | Yes    | Yes    |
| ItemGroup (Folder) | Yes    | Yes    |
| Item               | Yes    | Yes    |
| Request            | Yes    | Yes    |
| Response           | No     | No     |
| Event              | Yes    | Yes    |
| Variable           | Yes    | Yes    |
| Auth               | Yes    | Yes    |
