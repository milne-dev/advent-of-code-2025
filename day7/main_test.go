package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	lines := []string{}
	want := 1
	got := PartOne(lines)
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}

func TestPartTwo(t *testing.T) {
	lines := []string{}
	want := 1
	got := PartTwo(lines)
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}
