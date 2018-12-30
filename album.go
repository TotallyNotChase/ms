package main

// Album is just the name of each record to be scheduled.
type Album string

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
}

// RemoveAlbum gets rid of an instance of album on the given block.
func (b *Block) RemoveAlbum(album Album) {
	var newAlbums []Album

	for _, blockAlbum := range b.Albums {
		if album != blockAlbum {
			newAlbums = append(newAlbums, album)
		}
	}

	b.Albums = newAlbums
}

// ReplaceAlbum iterates through the list of albums and replaces an album
// with a new one.
func (b *Block) ReplaceAlbum(old, actual Album) {
}
