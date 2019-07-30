package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	qRef = Queue{
		&Block{"week 4", []Album{Album{Name: "nice"}, Album{Name: "nice"}}},
		&Block{"week 3", []Album{Album{Name: "placeholder"}}},
		&Block{"week 2", []Album{Album{Name: "fizz"}, Album{Name: "buzz"}, Album{Name: "fizbuzz"}}},
		&Block{"week 1", []Album{Album{Name: "foo"}, Album{Name: "bar"}, Album{Name: "foobar"}}},
	}
)

func TestAddBlock(t *testing.T) {
	tmp := qRef
	qLocal := &tmp

	b := &Block{"week 5", []Album{Album{Name: "darude - sandstorm"}}}

	q := &Queue{
		b,
		&Block{"week 4", []Album{Album{Name: "nice"}, Album{Name: "nice"}}},
		&Block{"week 3", []Album{Album{Name: "placeholder"}}},
		&Block{"week 2", []Album{Album{Name: "fizz"}, Album{Name: "buzz"}, Album{Name: "fizbuzz"}}},
	}

	qLocal.Add(b)

	if !cmp.Equal(qLocal, q) {
		t.Error("ms | Error with AddBlock: didn't add properly")

		fmt.Println("Reference:")
		qRef.ShowCurrent()

		fmt.Println("\nNew:")
		q.ShowCurrent()
	}

	q2 := &Queue{}
	q2.Add(b)
	q3 := &Queue{b}

	if !cmp.Equal(q2, q3) {
		t.Error("ms | Error with AddBlock: couldn't add properly to empty queue")
	}
}
