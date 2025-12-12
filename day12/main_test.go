package main

import (
	"aoc2025/utils"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`
	want := 2
	got := PartOne(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}
