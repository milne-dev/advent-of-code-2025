package main

import (
	"aoc2025/utils"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := ``
	want := 1
	got := PartOne(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}
