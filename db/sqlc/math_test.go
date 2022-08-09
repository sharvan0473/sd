package db

import (
	"github.com/sharvan/simplebank/api"
	"testing"
)

func TestAdd(t *testing.T) {

	got := api.Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
