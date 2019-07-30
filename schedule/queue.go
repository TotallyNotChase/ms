package schedule

import (
	"github.com/adrg/xdg"
	"gitlab.com/Sacules/jsonfile"
)

const (
	app = "/ms"
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

// Load queue from disk.
func (q *Queue) Load() error {
	err := createDirNotExist(xdg.ConfigHome + app)
	if err != nil {
		return err
	}

	return jsonfile.Load(q, xdg.ConfigHome+app+"/current.json")
}

// Replace goes through the queue and replaces an album for a new one.
func (q *Queue) Replace(old, actual Album) {
	for _, block := range q {
		block.Replace(old, actual)
	}
}

// Save queue to disk.
func (q *Queue) Save() error {
	err := createDirNotExist(xdg.ConfigHome + app)
	if err != nil {
		return err
	}

	return jsonfile.Save(q, xdg.ConfigHome+app+"/current.json")
}
