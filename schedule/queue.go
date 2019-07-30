package schedule

import (
	"path/filepath"

	"github.com/adrg/xdg"
	"gitlab.com/Sacules/jsonfile"
)

const (
	app      = "ms"
	savefile = "current.json"
)

var (
	dataDir  = filepath.Join(xdg.DataHome, app)
	dataPath = filepath.Join(dataDir, savefile)
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
	err := createDirNotExist(dataDir)
	if err != nil {
		return err
	}

	return jsonfile.Load(q, dataPath)
}

// Replace goes through the queue and replaces an album for a new one.
func (q *Queue) Replace(old, actual Album) {
	for _, block := range q {
		block.Replace(old, actual)
	}
}

// Save queue to disk.
func (q *Queue) Save() error {
	err := createDirNotExist(dataDir)
	if err != nil {
		return err
	}

	return jsonfile.Save(q, dataPath)
}
