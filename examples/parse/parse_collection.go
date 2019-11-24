package main

import (
	"fmt"

	"github.com/rbretecher/go-postman-collection"
)

func main() {

	c, err := postman.ParseCollection("examples/parse/parse_collection.json")

	if err != nil {
		panic(err)
	}

	exploreItems(c.Items)
}

func exploreItems(items []postman.Items) {
	for _, i := range items {
		if i.IsGroup() {
			folder := i.(*postman.ItemGroup)
			println(">", folder.Name)
			exploreItems(folder.Items)
		} else {
			item := i.(*postman.Item)
			if item.Request != nil {
				println(fmt.Sprintf("[%s] - %s - %s", item.Request.Method, item.Request.URL, item.Name))
			}
		}
	}
}
