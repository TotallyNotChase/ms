package main

import (
	"flag"
)

var (
	// Flags
	verbose bool
)

func main() {
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()

	// Testing
	api := InitAPI()
	LookupArtwork(api, "dark side of the moon", "cache/week 2", verbose)
}
