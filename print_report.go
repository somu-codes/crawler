package main

import (
	"fmt"
	"sort"
)

type pageCount struct {
	url   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	// Header
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	// Convert map to slice
	var sortedPages []pageCount
	for url, count := range pages {
		sortedPages = append(sortedPages, pageCount{url: url, count: count})
	}

	// Sort: highest count first, then alphabetically
	sort.Slice(sortedPages, func(i, j int) bool {
		if sortedPages[i].count == sortedPages[j].count {
			return sortedPages[i].url < sortedPages[j].url
		}
		return sortedPages[i].count > sortedPages[j].count
	})

	// Print in required format
	for _, p := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", p.count, p.url)
	}
}
