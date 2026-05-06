# mutating maps


```go
	websiteUrls := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	// Adding new item in maps
	websiteUrls["Facebook"] = "https://facebook.com"

  delete(websiteUrls, "Google")
	fmt.Println(websiteUrls)

```


Next: [Maps vs structs](./12-maps-vs-structs.md)