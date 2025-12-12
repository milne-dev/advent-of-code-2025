package main

import (
	"aoc2025/utils"
	"fmt"
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

func TestRotatePresent(t *testing.T) {
	cases := [][2]Present{
		{Present{
			{true, true, false},
			{true, true, false},
			{true, true, false},
		}, Present{
			{true, true, true},
			{true, true, true},
			{false, false, false},
		}},
		{Present{
			{true, true, true},
			{true, true, true},
			{false, false, false},
		}, Present{
			{false, true, true},
			{false, true, true},
			{false, true, true},
		}},
		{Present{
			{false, true, true},
			{false, true, true},
			{false, true, true},
		}, Present{
			{false, false, false},
			{true, true, true},
			{true, true, true},
		}},
		{Present{
			{false, false, false},
			{true, true, true},
			{true, true, true},
		}, Present{
			{true, true, false},
			{true, true, false},
			{true, true, false},
		}},
	}
	for _, testCase := range cases {
		got := testCase[0]
		RotatePresent(got)
		want := testCase[1]
		if want != got {
			t.Errorf("RotatePresent(...)\n")
			fmt.Println("WANT:")
			for i := range want {
				for j := range want[i] {
					fmt.Printf("%v", boolToSym(want[i][j]))
				}
				fmt.Println()
			}
			fmt.Println("GOT:")
			for i := range got {
				for j := range got[i] {
					fmt.Printf("%v", boolToSym(got[i][j]))
				}
				fmt.Println()
			}
		}
	}
}
