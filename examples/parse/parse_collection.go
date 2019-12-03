package main

import (
	"fmt"
	"os"

	"github.com/rbretecher/go-postman-collection"
)

func main() {

	file, err := os.Open("testdata/basic_collection.json")
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

func exploreItems(items []*postman.Items) {
	for _, i := range items {
		if i.IsGroup() {
			println(">", i.Name)
			exploreItems(i.Items)
		} else {
			if i.Request != nil {
				println(fmt.Sprintf("[%s] - %s - %s", i.Request.Method, i.Request.URL, i.Name))

				if i.Request.Auth != nil {
					for _, p := range i.Request.Auth.GetParams() {
						println(p.Key, ":", fmt.Sprintf("%v", p.Value))
					}
				}
			} else {
				println(i.Name)
			}
		}
	}
}
