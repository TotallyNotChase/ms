package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	name = "nice"
	foo  = NewAlbum("foo")
	bar  = NewAlbum("bar")
	al   = []Album{foo, bar}
	bRef = Block{name, al}
)

func TestNewBlock(t *testing.T) {
	bNew := NewBlock(name, foo, bar)

	if bNew == nil {
		t.Error("ms | Error with NewBlock: nil pointer")
	}

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
		t.Error("ms | Error with RemoveAlbum: not working on empty block")
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
		t.Error("ms | Error with ReplaceAlbum: generated struct not equal to testing one")
	}
}
