package main

import (
	"reflect"
	"testing"
)

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

func TestAdvanceRunes(t *testing.T) {
	var got, want []rune

	got = advanceRunes([]rune{'Z'})
	want = []rune{'0', '0'}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error: got runes %s, want %s.", string(got), string(want))
	}

	got = advanceRunes([]rune{'0', '0'})
	want = []rune{'0', '1'}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error: got runes %s, want %s.", string(got), string(want))
	}

	got = advanceRunes([]rune{'9', '9'})
	want = []rune{'9', 'a'}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error: got runes %s, want %s.", string(got), string(want))
	}

	got = advanceRunes([]rune{'Z', 'Z'})
	want = []rune{'0', '0', '0'}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error: got runes %s, want %s.", string(got), string(want))
	}
}
