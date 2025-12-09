package main

import (
	"aoc2025/utils"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	want := 50
	got := PartOne(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	want := 24
	got := PartTwo(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}

	input = `7,1
11,1
11,6
12,6
12,1
14,1
14,7
9,7
9,5
2,5
2,3
7,3`
	want = 42
	got = PartTwo(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}

	input = `7,1
11,1
11,6
13,6
13,1
14,1
14,7
9,7
9,5
2,5
2,3
7,3`
	want = 24
	got = PartTwo(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}
