package main

import "testing"

func TestPartOne(t *testing.T) {
	d := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
		"565653-565659",
		"824824821-824824827",
		"2121212118-2121212124",
	}
	want := 1227775554
	got := PartOne(d)
	if want != got {
		t.Errorf(`PartOne() want %v, got %v`, want, got)
	}
}

func TestIsInvalid(t *testing.T) {
	d := []struct {
		id  string
		res bool
	}{
		{"22", true},
		{"21", false},
		{"101", false},
	}

	for _, c := range d {
		want := c.res
		got := IsInvalid(c.id)
		if want != got {
			t.Errorf(`IsInvalid(%v) want %v, got %v`, c.id,
				want, got)
		}
	}
}
