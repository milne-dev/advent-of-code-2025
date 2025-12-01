package main

import "testing"

func TestSolve(t *testing.T) {
	d := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	want := 3
	got := Solve(d)
	if want != got {
		t.Errorf(`Solve() want %v, got %v`, want, got)
	}
}

func TestChangePos(t *testing.T) {
	cases := [][]int{
		{50, 50, 0},
		{50, -50, 0},
		{50, 51, 1},
		{50, -51, 99},
		{50, -68, 82},
		{99, 100, 99},
		{99, 500, 99},
		{99, 550, 49},
		{99, 1000, 99},
		{0, -1, 99},
		{0, -100, 0},
		{0, -500, 0},
		{0, -1000, 0},
		{82, -30, 52},
		{55, -55, 0},
	}
	for _, c := range cases {
		want := c[2]
		got := ChangePos(c[0], c[1])
		if want != got {
			t.Errorf(`ChangePos(%v, %v) want %v, got %v`, c[0], c[1],
				want, got)
		}
	}
}
