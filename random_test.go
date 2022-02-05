package main

import (
	"testing"
)

func TestDistrictName(t *testing.T) {
	t.Run("引数のnumに一致する区名を返す(足立区)", func(t *testing.T) {
		got := DistrictName(1)
		want := "足立区"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("引数のnumに一致する区名を返す(足立区)", func(t *testing.T) {
		got := DistrictName(12)
		want := "杉並区"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
