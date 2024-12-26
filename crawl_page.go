package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()

	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse current URL: %v\n", err)
		return
	}

	// if the URL is not from the same domain, dont crawl the page to avoid crawling the whole internet
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizarUrl(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't normalize URL: %v\n", err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	// if the URL is already in the map, dont crawl the page just increment the count
	if !isFirst {
		return
	}

	fmt.Printf("Crawling: %s\n", rawCurrentURL)
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't get HTML: %v\n", err)
		return
	}
	fmt.Println("Got HTML")

	nextUrls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't get URLs from HTML: %v\n", err)
		return
	}
	// recursively crawl the tree of URLs
	for _, url := range nextUrls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
