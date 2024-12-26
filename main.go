package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Missing arguments, usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error - maxConcurrency parameter should be a number: %v\n", err)
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error - maxPages parameter should be a number: %v\n", err)
		os.Exit(1)
	}

	cfg, err := configure(baseUrl, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Starting crawl of: %s \n", baseUrl)
	// Timer for meassuring the time it takes to crawl the website
	start := time.Now()
	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	elapsed := time.Since(start)

	fmt.Println("Results: ")
	for url, count := range cfg.pages {
		fmt.Printf("%s: %d\n", url, count)
	}
	fmt.Println("Crawl completed - Elapsed time: ", elapsed)

	os.Exit(0)
}
