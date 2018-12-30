package main

// Queue is the overall scheduler, consisting of 4 blocks.
type Queue struct {
	Block [4]*Block `json:"block"`
}

// Add a block to the queue and get rid of any old ones.
func (q *Queue) Add(b *Block) {
}

// Replace goes through the queue and replaces an album for a new one.
func (q *Queue) Replace(old, actual Album) {
}

// ShowCurrent prints the current week of records in the queue.
func (q *Queue) ShowCurrent() {
}
