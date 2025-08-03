package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: go run main.go <url> [maxConcurrency] [maxPages]")
		os.Exit(1)
	}

	baseURL := args[0]

	// Default values
	maxConcurrency := 5
	maxPages := 50

	// Parse optional args
	if len(args) >= 2 {
		if mc, err := strconv.Atoi(args[1]); err == nil {
			maxConcurrency = mc
		}
	}
	if len(args) >= 3 {
		if mp, err := strconv.Atoi(args[2]); err == nil {
			maxPages = mp
		}
	}

	fmt.Printf("starting crawl of: %s\n", baseURL)
	fmt.Printf("Max concurrency: %d\n", maxConcurrency)
	fmt.Printf("Max pages: %d\n", maxPages)

	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("invalid base URL:", err)
		os.Exit(1)
	}

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            parsedBase,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, baseURL)
}
