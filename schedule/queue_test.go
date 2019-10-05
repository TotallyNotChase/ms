package schedule

import (
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
		t.Error("didn't add properly")
	}

	q2 := &Queue{}
	q2.Add(b)
	q3 := &Queue{b}

	if !cmp.Equal(q2, q3) {
		t.Error("couldn't add properly to empty queue")
	}
}

func TestReplaceAlbumInQueue(t *testing.T) {

	q1 := &Queue{
		&Block{"week 4", []Album{Album{Name: "nice"}, Album{Name: "nice"}}},
		&Block{"week 3", []Album{Album{Name: "placeholder"}}},
		&Block{"week 2", []Album{Album{Name: "fizz"}, Album{Name: "buzz"}, Album{Name: "fizbuzz"}}},
		&Block{"week 1", []Album{Album{Name: "foo"}, Album{Name: "bar"}, Album{Name: "foobar"}}},
	}

	q2 := &Queue{
		&Block{"week 4", []Album{Album{Name: "nice"}, Album{Name: "nice"}}},
		&Block{"week 3", []Album{Album{Name: "placeholder"}}},
		&Block{"week 2", []Album{Album{Name: "fuzz"}, Album{Name: "buzz"}, Album{Name: "fizbuzz"}}},
		&Block{"week 1", []Album{Album{Name: "foo"}, Album{Name: "bar"}, Album{Name: "foobar"}}},
	}

	oldAlbum := Album{Name: "fizz"}
	newAlbum := Album{Name: "fuzz"}
	q1.Replace(oldAlbum, newAlbum)

	if !cmp.Equal(q1, q2) {
		t.Error("couldn't replace an album in queue")
	}
}
