package main

import "testing"

func TestTest(t *testing.T) {
	got := 1 + 1
	want := 2

	if got != want {
		t.Error("Can't happen")
	}
}
