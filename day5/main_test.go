package main

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	}
	want := 3
	got := PartOne(input)
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}
