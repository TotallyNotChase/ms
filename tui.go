package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type tui struct {
	app     *tview.Application
	flex    *tview.Flex
	blocks  []*tview.Table
	headers []*tview.TableCell
	queue   *Queue
}

func newTui() *tui {
	return &tui{
		app:    tview.NewApplication(),
		flex:   tview.NewFlex(),
		blocks: make([]*tview.Table, 0),
	}
}

func (tui *tui) addBlocks() {
	// Queue
	for i, block := range tui.queue {
		// Skip resting week
		if i == 2 {
			continue
		}

		// Each block
		blocktable := tview.NewTable()
		blocktable.SetTitle("[::b] " + block.Name + " ")

		// Headers
		albumcell := tview.NewTableCell("Albums")
		listenedcell := tview.NewTableCell("Listened")
		ratedcell := tview.NewTableCell("Rated")

		// Add headers
		blocktable.SetCell(0, 0, albumcell)
		blocktable.SetCell(0, 1, listenedcell)
		blocktable.SetCell(0, 2, ratedcell)
		tui.headers = append(tui.headers, albumcell)
		tui.headers = append(tui.headers, listenedcell)
		tui.headers = append(tui.headers, ratedcell)

		// Albums
		for j, album := range block.Albums {
			var listened, rated string

			cell := tview.NewTableCell(album.Name)
			cell.
				SetExpansion(1).
				SetMaxWidth(len(album.Name))

			blocktable.SetCell(j+1, 0, cell)

			// Check per block
			if i == 0 && album.FirstListen {
				listened = "✓"
			}
			if i == 1 && album.SecondListen {
				listened = "✓"
			}
			if i == 3 && album.ThirdListen {
				listened = "✓"
			}
			if album.Rated {
				rated = "✓"
			}

			blocktable.SetCell(j+1, 1, tview.NewTableCell(listened).SetAlign(tview.AlignCenter))
			blocktable.SetCell(j+1, 2, tview.NewTableCell(rated).SetAlign(tview.AlignCenter))
		}

		tui.flex.AddItem(blocktable, 0, 1, true)
		tui.blocks = append(tui.blocks, blocktable)
	}
}

func (tui *tui) setupLooks() {
	// Borders
	tview.Borders.HorizontalFocus = tview.BoxDrawingsHeavyHorizontal
	tview.Borders.VerticalFocus = tview.BoxDrawingsHeavyVertical
	tview.Borders.TopLeftFocus = tview.BoxDrawingsHeavyDownAndRight
	tview.Borders.TopRightFocus = tview.BoxDrawingsHeavyDownAndLeft
	tview.Borders.BottomLeftFocus = tview.BoxDrawingsHeavyUpAndRight
	tview.Borders.BottomRightFocus = tview.BoxDrawingsHeavyUpAndLeft

	// Flex
	tui.flex.
		SetDirection(tview.FlexRow)

	// Blocks
	for _, blocktable := range tui.blocks {
		blocktable.
			SetSelectable(true, true).
			SetFixed(1, 3).
			SetTitleColor(tcell.ColorGreen).
			SetBorder(true)
	}

	// Block headers
	for _, header := range tui.headers {
		header.
			SetSelectable(false).
			SetTextColor(tcell.ColorYellow).
			SetAttributes(tcell.AttrBold)

	}
}

func (tui *tui) setupBindings() {
	// Main binds
	tui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			tui.app.Stop()
		}

		return event
	})

	// Per-block binds
	for i, blocktable := range tui.blocks {
		blocktable.SetSelectedFunc(func(row, column int) {
			cell := blocktable.GetCell(row, column)

			switch column {
			// Listened
			case 1:
				if cell.Text == "" {
					cell.SetText("✓")
				} else {
					cell.SetText("")
				}

			// Rated
			case 2:
				if cell.Text == "" {
					cell.SetText("✓")
					tui.queue[i].Albums[row-1].Rated = true
				} else {
					cell.SetText("")
					tui.queue[i].Albums[row-1].Rated = false
				}
			}
		})
	}
}

func (tui *tui) run() {
	// Run the App
	tui.app.SetRoot(tui.flex, true).SetFocus(tui.flex)
	if err := tui.app.Run(); err != nil {
		panic(err)
	}
}

func (tui *tui) init() {
	var q = NewQueue()

	if err := q.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tui.queue = q

	tui.addBlocks()
	tui.setupLooks()
	tui.setupBindings()
}
