package box

import (
	"context"
	"testing"
)

var keyA = boxkey{"A"}
var keyB = boxkey{"B"}
var keyC = boxkey{"C"}

func TestSelf(t *testing.T) {
	box := WithValue(context.Background(), keyA, "a")

	box1 := WithValue(box, keyB, "b")

	if box1 != box {
		t.Fatalf("box1 != box")
	}
}

func TestNil(t *testing.T) {
	box := WithValue(context.Background(), keyA, "a")

	c, _ := From[string](box, keyC)
	if c != "" {
		t.Fatalf("c is not empty")
	}
}
