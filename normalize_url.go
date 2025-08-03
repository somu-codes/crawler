package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(inputURL string) (outputUrl string, err error) {
	parsed, err := url.Parse(inputURL)
	if err != nil {
		return
	}
	fmt.Printf(parsed.Host + " " + parsed.Path)
	outputUrl = parsed.Host + parsed.Path
	fmt.Println(outputUrl) // blog.boot.dev/path
	return
}
