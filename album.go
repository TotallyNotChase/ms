package main

// Album is a record to be scheduled. Includes whether it's been listened
// or rated.
type Album struct {
	Name     string `json:"name"`
	Listened bool   `json:"listened"`
	Rated    bool   `json:"rated"`
}

// NewAlbum returns a pointer to an empty Album.
func NewAlbum(al string) *Album {
	return &Album{Name: al}
}
