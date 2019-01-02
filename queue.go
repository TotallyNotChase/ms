package main

import (
	"fmt"

	"github.com/adrg/xdg"
	"gitlab.com/Sacules/jsonfile"
)

const (
	app = "/ms"
)

// Queue is the overall scheduler, consisting of 4 blocks.
type Queue [4]*Block

// NewQueue returns a pointer to an empty Queue.
func NewQueue() *Queue {
	return &Queue{}
}

// Add a block to the queue and get rid of any old ones.
func (q *Queue) Add(b *Block) {
	for i := len(q) - 1; i > 0; i-- {
		q[i] = q[i-1]
	}

	q[0] = b
}

// Load queue from disk.
func (q *Queue) Load() error {
	err := createDirNotExist(xdg.ConfigHome + app)
	if err != nil {
		return err
	}

	return jsonfile.LoadFile(q, xdg.ConfigHome+app+"/current.json")
}

// Replace goes through the queue and replaces an album for a new one.
func (q *Queue) Replace(old, actual Album) {
	for _, block := range q {
		block.ReplaceAlbum(old, actual)
	}
}

// Save queue to disk.
func (q *Queue) Save() error {
	err := createDirNotExist(xdg.ConfigHome + app)
	if err != nil {
		return err
	}

	return jsonfile.SaveFile(q, xdg.ConfigHome+app+"/current.json")
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
