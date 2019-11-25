package main

import (
	"fmt"
	"os"

	"github.com/rbretecher/go-postman-collection"
)

func main() {

	file, err := os.Open("examples/parse/parse_collection.json")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	c, err := postman.ParseCollection(file)

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

				if item.Request.Auth != nil {
					for _, p := range item.Request.Auth.GetParams() {
						println(p.Key, ":", fmt.Sprintf("%v", p.Value))
					}
				}
			}
		}
	}
}
