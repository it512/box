package box

import (
	"context"
	"testing"
)

type BoxKey struct {
	Name string
}

func TestSelf(t *testing.T) {
	b := New(context.Background())
	b.Add("c", "c")
	x := context.WithValue(b, BoxKey{"A"}, "b")

	if Self(x) != b {
		t.Fail()
	}

	Self(x).Add("z", "z")

	if From[string](x, "c") != "c" {
		t.Fail()
	}

	if From[string](x, "z") != "z" {
		t.Fail()
	}
}
