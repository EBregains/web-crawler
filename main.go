package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) >= 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := args[1]
	fmt.Printf("starting crawl of: %s", baseUrl)
	os.Exit(0)
}
