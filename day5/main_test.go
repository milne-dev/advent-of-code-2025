package main

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{}
	want := 0
	got := PartOne(input)
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}
