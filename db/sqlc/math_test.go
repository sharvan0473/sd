package db

import (
	"testing"
)

func Add(x, y int) (res int) {
	return x + y
}

func TestAdd(t *testing.T) {

	got := Add(4, 4)
	want := 8

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
