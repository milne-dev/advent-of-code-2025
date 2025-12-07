package main

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{}
	want := 3
	got := PartOne(input)
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{}
	want := 14
	got := PartTwo(input)
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}
