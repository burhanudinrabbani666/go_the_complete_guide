package main

import "fmt"

func main() {

	websiteUrls := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	// Adding new item in maps
	websiteUrls["Facebook"] = "https://facebook.com"

	delete(websiteUrls, "Google")
	fmt.Println(websiteUrls)

}
