# go-postman-collection

Go module to work with Postman Collections. It only works with Postman Collection format v2.1.0.

### Example

```go
// Create an empty collection
c := CreateCollection("Test", "Awesome description")
c.Write("collection.json")
```

### Useful resources

https://schema.getpostman.com/json/collection/latest/docs/index.html
