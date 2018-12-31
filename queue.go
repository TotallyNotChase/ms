package main

import (
	"fmt"
	"gitlab.com/Sacules/jsonfile"
)

// Queue is the overall scheduler, consisting of 4 blocks.
type Queue [4]*Block

// Add a block to the queue and get rid of any old ones.
func (q *Queue) Add(b *Block) {
	for i := len(q) - 1; i > 0; i-- {
		q[i] = q[i-1]
	}

	q[0] = b
}

// Replace goes through the queue and replaces an album for a new one.
func (q *Queue) Replace(old, actual Album) {
	for _, block := range q {
		block.ReplaceAlbum(old, actual)
	}
}

// ShowCurrent prints the current week of records in the queue.
func (q *Queue) ShowCurrent() {
	for _, block := range q {
		if block != nil {
			fmt.Printf("\n%s\n", block.Name)
			for i, album := range block.Albums {
				fmt.Printf("%d. %s\n", i+1, album)
			}
		}
	}
}

// Load queue from disk.
func (q *Queue) Load() error {
	return jsonfile.LoadFile(q, "queue.json")
}

// Save queue to disk.
func (q *Queue) Save() error {
	return jsonfile.SaveFile(q, "queue.json")
}
