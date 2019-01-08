package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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
		newblock()

	case "newqueue":
	default:
		flag.PrintDefaults()
	}

	// Testing
	// api := InitAPI()
	// LookupArtwork(api, "dark side of the moon", "cache/week 2", verbose)
}

func status() {
	var q = NewQueue()
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

func newblock() {
	var (
		scanner = bufio.NewScanner(os.Stdin)
		q       = NewQueue()
	)

	err := q.Load()
	if err != nil {
		fmt.Printf("Error reading queue: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Name for new block: ")
	scanner.Scan()
	name := scanner.Text()

records:
	fmt.Printf("Amount of records for this week: ")
	scanner.Scan()
	num := scanner.Text()
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Printf("Error with amount of record: %s\n", err)
		fmt.Println("Pleae try again.")
		goto records
	}
	albums := make([]Album, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Album %d: ", i+1)
		scanner.Scan()
		al := scanner.Text()
		if al == "" {
			break
		}

		albums[i].Name = al
	}

	b := NewBlock(name, albums...)
	q.Add(b)

	err = q.Save()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
