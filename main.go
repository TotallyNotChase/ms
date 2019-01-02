package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shkh/lastfm-go/lastfm"
)

var (
	// Flags
	verbose bool
)

func main() {
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

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
	LookupArtwork(api, "dark side of the moon", "cache/week 2", verbose)
}
