package main

import (
	"os"
  "os/signal"
)

func main() {
	var test string
	exitsig := make(chan os.Signal, 1)
	signal.Notify(exitsig, os.Interrupt)
	var (
			lockstate bool = false
	)

	go func() {
			<-exitsig
			if lockstate {
					var err = os.Remove("ms.lock")
					if err != nil {
							return
					}
			}
			os.Exit(0)
	}()

	if _, err := os.Stat("ms.lock"); err == nil {
			return

	} else if os.IsNotExist(err) {
			var file, err = os.Create("ms.lock")
			if err != nil {
					return
			}
			file.Close()
			defer os.Remove("ms.lock")
			lockstate = true
	}
	
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
