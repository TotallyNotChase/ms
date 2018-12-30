package main

// Queue is the overall scheduler, consisting of 4 blocks.
type Queue struct {
	Block [4]*Block `json:"block"`
}

// AddBlock appends a block to the queue and gets rid of any old ones.
func AddBlock(q Queue) {
}

// ReplaceInBlock goes through the given queue and replaces an album for a new one.
func ReplaceInBlock(q Queue, old, actual Album) {
}

// ShowCurrent prints the current week of records in the queue.
func ShowCurrent(q Queue) {
}
