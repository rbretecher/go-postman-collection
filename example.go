package main

func main() {
	c := CreateCollection("Go Collection", "Awesome description")

	c.AddFolder("This is a folder").AddItem(&Item{
		Name: "An item inside a folder",
	})

	c.AddItem(&Item{
		Name: "This is a request",
	})

	c.AddFolder("Empty folder")

	c.Write("postman_collection.json")
}
