package main

import (
	"testing"
)

func TestDistrictName(t *testing.T) {
	got := DistrictName(1)
	want := "足立区"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
