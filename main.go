package main

import (
	"os"
)

var (
	// Flags
	verbose bool
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "newblock" {
			newblock()
		}

		return
	}

	tui := newTui()

	tui.init()
	tui.run()
}
