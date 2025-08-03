package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

// getURLsFromHTML extracts all links from HTML and converts relative URLs to absolute.
func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var urls []string
	var visitNode func(*html.Node)

	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := strings.TrimSpace(attr.Val)
					if href == "" {
						continue
					}

					// Convert relative URL to absolute
					u, err := url.Parse(href)
					if err == nil {
						resolved := base.ResolveReference(u)
						urls = append(urls, resolved.String())
					}
				}
			}
		}
		// Recursively visit child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	visitNode(doc)
	return urls, nil
}
