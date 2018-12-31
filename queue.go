package main

import (
	"fmt"
)

// Queue is the overall scheduler, consisting of 4 blocks.
type Queue [4]*Block

// type Queue struct {
// 	Blocks [4]*Block `json:"block"`
// }

// Add a block to the queue and get rid of any old ones.
func (q *Queue) Add(b *Block) {
	for i := len(q); i > 0; i-- {
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
		fmt.Printf("\n%s", block.Name)
		for i, album := range block.Albums {
			fmt.Printf("%d. %s", i+1, album)
		}
	}
}
