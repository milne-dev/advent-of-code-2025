package main

import "testing"

func TestPartOne(t *testing.T) {
	input := ""
	want := 357
	got := PartOne(input)
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}
