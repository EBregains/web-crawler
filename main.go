package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := os.Args[1]
	fmt.Printf("starting crawl of: %s", baseUrl)

	pages := make(map[string]int)

	// Timer for meassuring the time it takes to crawl the website
	start := time.Now()
	crawlPage(baseUrl, baseUrl, pages)
	elapsed := time.Since(start)

	fmt.Println("Results: ")
	for url, count := range pages {
		fmt.Printf("%s: %d\n", url, count)
	}
	fmt.Println("Crawl completed - Elapsed time: ", elapsed)

	os.Exit(0)
}
