package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	url := flag.String("url", "", "URL not null")
	concurrencyStr := flag.String("concurrency", "50", "concurrency not null")

	flag.Parse()

	if *url == "" {
		fmt.Println("url must not null")
	}

	concurrency, _ := strconv.Atoi(*concurrencyStr)
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			resp, err := http.Get(*url)
			if err != nil {
				fmt.Printf("Error send Get reuqest: %s\n", err.Error())
			}
			defer resp.Body.Close()
			fmt.Println("Request completed")
		}()
	}
	wg.Wait()
	fmt.Println("All request completed")
}
