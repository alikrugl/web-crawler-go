package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("waiting for 3 arguments")
		return
	}

	rawBaseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error - invalid maxConcurrency: %v", err)
		return
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error - invalid maxPages: %v", err)
		return
	}
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
