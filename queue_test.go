package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	qRef = Queue{
		NewBlock("week 4", "nice", "nice"),
		NewBlock("week 3", "placeholder"),
		NewBlock("week 2", "fizz", "buzz", "fizbuzz"),
		NewBlock("week 1", "foo", "bar", "foobar"),
	}
)

func TestAddBlock(t *testing.T) {
	tmp := qRef
	qLocal := &tmp

	b := NewBlock("week 5", "darude - sandstorm")

	q := &Queue{
		b,
		NewBlock("week 4", "nice", "nice"),
		NewBlock("week 3", "placeholder"),
		NewBlock("week 2", "fizz", "buzz", "fizbuzz"),
	}

	qLocal.Add(b)

	if !cmp.Equal(qLocal, q) {
		t.Error("ms | Error with AddBlock: didn't add properly")

		fmt.Println("Reference:")
		qRef.ShowCurrent()

		fmt.Println("\nNew:")
		q.ShowCurrent()
	}
}
