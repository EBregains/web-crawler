package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse base URL: %v\n", err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse current URL: %v\n", err)
		return
	}
	if currentURL.Host != baseURL.Host {
		return
	}

	normalizedURL, err := normalizarUrl(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't normalize URL: %v\n", err)
		return
	}
	// if the URL is already in the map, dont crawl the page just increment the count
	if _, ok := pages[normalizedURL]; ok {
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL] = 1
	fmt.Printf("Crawling: %s\n", rawCurrentURL)
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't get HTML: %v\n", err)
		return
	}
	fmt.Printf("Got HTML for: %s\n", rawCurrentURL)
	nextUrls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't get URLs from HTML: %v\n", err)
		return
	}
	// recursively crawl the tree of URLs
	for _, url := range nextUrls {
		crawlPage(rawBaseURL, url, pages)
	}
}
