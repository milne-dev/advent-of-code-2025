package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `
	lines := strings.Lines(input)
	want := 4277556
	got := PartOne(lines)
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}

	input = `1 1 1 1 1
		1 1 1 1 1
		1 1 1 1 1
		+ + + + +`
	lines = strings.Lines(input)
	want = 15
	got = PartOne(lines)
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `
	lines := strings.Lines(input)
	want := 3263827
	got := PartTwo(lines)
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}
