package main

import "testing"

func TestGetNextRune(t *testing.T) {
	var got, want rune

	got = getNextRune('9')
	want = 'a'
	if got != want {
		t.Errorf("error: got rune %s, want %s.", string(got), string(want))
	}

	got = getNextRune('z')
	want = 'A'
	if got != want {
		t.Errorf("error: got rune %s, want %s.", string(got), string(want))
	}

	got = getNextRune('Z')
	want = '0'
	if got != want {
		t.Errorf("error: got rune %s, want %s.", string(got), string(want))
	}
}
