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

	htmlBody, err := getHTML(baseUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(htmlBody)
	os.Exit(0)
}
