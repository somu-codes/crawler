package main

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if count, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL] = count + 1
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

	// Enforce concurrency limit
	cfg.concurrencyControl <- struct{}{}        // acquire slot
	defer func() { <-cfg.concurrencyControl }() // release slot

	// Stop if we've reached the maxPages limit
	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	// Stay on same domain
	if !strings.HasPrefix(rawCurrentURL, cfg.baseURL.String()) {
		return
	}

	// Normalize URL
	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Error normalizing:", rawCurrentURL, err)
		return
	}

	// Check if we already crawled this page
	if !cfg.addPageVisit(normURL) {
		return
	}

	// Fetch HTML
	fmt.Println("Crawling:", rawCurrentURL)
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error fetching:", rawCurrentURL, err)
		return
	}

	// Extract internal links
	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Println("Error extracting URLs from:", rawCurrentURL, err)
		return
	}

	// Crawl all found URLs concurrently
	for _, u := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}
}
