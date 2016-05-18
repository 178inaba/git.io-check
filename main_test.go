package main

import "testing"

func TestGetNextRune(t *testing.T) {
	var got, want rune

	got = getNextRune('a')
	want = 'b'
	if got != want {
		t.Errorf("error: got rune %s, want %s.", string(got), string(want))
	}
}
