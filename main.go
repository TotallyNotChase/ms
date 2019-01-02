package main

import (
	"flag"
	"os"
)

var (
	// Flags
	verbose bool
)

func main() {
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "status":
		q := &Queue{}
		q.Load()
		q.ShowCurrent()

	default:
		flag.PrintDefaults()
	}

	// Testing
	// api := InitAPI()
	// LookupArtwork(api, "dark side of the moon", "cache/week 2", verbose)
}
