package schedule

// Block represents a week worth of albums, each week having a new block.
type Block struct {
	Name   string  `json:"name"`
	Albums []Album `json:"albums"`
}

// Add an album to the block.
func (b *Block) Add(al Album) {
	b.Albums = append(b.Albums, al)
}

// Remove gets rid of an instance of album on the given block.
func (b *Block) Remove(al Album) {
	var newAlbums []Album

	for _, blockAl := range b.Albums {
		if al != blockAl {
			newAlbums = append(newAlbums, blockAl)
		}
	}

	b.Albums = newAlbums
}

// Replace iterates through the list of albums and replaces an album
// with a new one.
func (b *Block) Replace(old, actual Album) {
	for i, al := range b.Albums {
		if al == old {
			b.Albums[i] = actual
			break
		}
	}
}
