package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/shkh/lastfm-go/lastfm"
)

func main() {
	apiKey := &APIKey{}
	err := apiKey.Load()
	if err != nil {
		fmt.Println("API Key couldn't be found, creating a new apikey.json file.")
		fmt.Println("Please put your keys there before using this program.")
		apiKey.Save()
		os.Exit(1)
	}

	api := lastfm.New(apiKey.Key, apiKey.Secret)

	// Testing
	var wg sync.WaitGroup
	results := make(chan string)
	errs := make(chan error)

	go func() {
		for {
			select {
			case res := <-results:
				fmt.Println(res)
			case err := <-errs:
				fmt.Println(err)
			}
		}
	}()

	GetArtwork(api, "thriller", "cache/week 1", results, errs, &wg)

	wg.Wait()
	close(results)
	close(errs)
}
