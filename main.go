package main

import (
	"fmt"
	"os"
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

	// htmlBody, err := getHTML(baseUrl)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	pages := make(map[string]int)
	crawlPage(baseUrl, baseUrl, pages)

	fmt.Println("Crawl complete")
	for url, count := range pages {
		fmt.Printf("%s: %d\n", url, count)
	}

	os.Exit(0)
}
