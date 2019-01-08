package main

// Block represents a week worth of albums, each week having a new block.
type Block struct {
	Name   string  `json:"name"`
	Albums []Album `json:"albums"`
}

// NewBlock inits a block with a name, usually the current week, and the
// albums it should include.
func NewBlock(name string, albums ...Album) *Block {
	return &Block{
		Name:   name,
		Albums: albums,
	}
}

// Add an album to the block.
func (b *Block) Add(al Album) {
	b.Albums = append(b.Albums, al)
}

// RemoveAlbum gets rid of an instance of album on the given block.
func (b *Block) RemoveAlbum(al Album) {
	var newAlbums []Album

	for _, blockAl := range b.Albums {
		if al != blockAl {
			newAlbums = append(newAlbums, blockAl)
		}
	}

	b.Albums = newAlbums
}

// ReplaceAlbum iterates through the list of albums and replaces an album
// with a new one.
func (b *Block) ReplaceAlbum(old, actual Album) {
	for i, al := range b.Albums {
		if al == old {
			b.Albums[i] = actual
			break
		}
	}
}
