package db

import (
	"testing"
)

func Add(x, y int) (res int) {
	return x + y
}

func TestAdd(t *testing.T) {

	got := Add(5, 5)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
