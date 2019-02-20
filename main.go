package main

var (
	// Flags
	verbose bool
)

func main() {
	tui := newTui()

	tui.init()
	tui.run()
}
