package main

import (
	"flag"
	"fmt"
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
		status()

	case "newblock":
	case "newqueue":
	default:
		flag.PrintDefaults()
	}

	// Testing
	// api := InitAPI()
	// LookupArtwork(api, "dark side of the moon", "cache/week 2", verbose)
}

func status() {
	q := &Queue{}
	err := q.Load()
	if err != nil {
		if q.Save() != nil {
			fmt.Println(err)
		}
		fmt.Println(err)
		os.Exit(1)
	}

	q.ShowCurrent()
}
