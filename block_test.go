package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	name = "nice"
	foo  = Album{Name: "foo"}
	bar  = Album{Name: "bar"}
	al   = []Album{foo, bar}
	bRef = Block{name, al}
)

func TestRemoveAlbum(t *testing.T) {
	tmp := bRef
	b := &tmp

	bLess := &Block{
		Name:   name,
		Albums: []Album{foo},
	}

	b.RemoveAlbum(bar)

	if !cmp.Equal(b, bLess) {
		t.Error("remove album: generated struct not equal to testing one")

		t.Error("Reference:")
		for i, album := range bLess.Albums {
			t.Errorf("\t%d. %s\n", i+1, album.Name)
		}

		t.Error("New one:")
		for i, album := range b.Albums {
			t.Errorf("\t%d. %s\n", i+1, album.Name)
		}
	}

	bEmpty := &Block{}
	bEmpty.RemoveAlbum(foo)

	if !cmp.Equal(bEmpty, &Block{}) {
		t.Error("remove album: not working on empty block")
	}
}

func TestReplaceAlbum(t *testing.T) {
	tmp := bRef
	b := &tmp

	bReplaced := &Block{
		Name:   name,
		Albums: []Album{foo, foo},
	}

	b.ReplaceAlbum(bar, foo)

	if !cmp.Equal(b, bReplaced) {
		t.Error("replac albumm: generated struct not equal to testing one")
	}
}
