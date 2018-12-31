package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewBlock(t *testing.T) {
	var (
		name       = "nice"
		foo  Album = "foo"
		bar  Album = "bar"
		al         = []Album{foo, bar}
	)

	bRef := &Block{name, al}

	bNew := NewBlock(name, foo, bar)

	if !cmp.Equal(bRef, bNew) {
		t.Error("ms | Error with NewBlock: generated struct not equal to testing one")
	}
}
