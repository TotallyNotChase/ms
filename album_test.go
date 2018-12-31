package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	name       = "nice"
	foo  Album = "foo"
	bar  Album = "bar"
	al         = []Album{foo, bar}
	bRef       = Block{name, al}
)

func TestNewBlock(t *testing.T) {
	bNew := NewBlock(name, foo, bar)

	if !cmp.Equal(&bRef, bNew) {
		t.Error("ms | Error with NewBlock: generated struct not equal to testing one")
	}
}

func TestRemoveAlbum(t *testing.T) {
	tmp := bRef
	b := &tmp

	bLess := &Block{
		Name:   name,
		Albums: []Album{foo},
	}

	b.RemoveAlbum(bar)

	if !cmp.Equal(b, bLess) {
		t.Error("ms | Error with RemoveAlbum: generated struct not equal to testing one")

		t.Error("Reference:")
		for i, album := range bLess.Albums {
			t.Errorf("\t%d. %s\n", i+1, album)
		}

		t.Error("New one:")
		for i, album := range b.Albums {
			t.Errorf("\t%d. %s\n", i+1, album)
		}
	}
}
