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
	fields := strings.Fields(input)
	want := 4277556
	got := PartOne(fields)
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
