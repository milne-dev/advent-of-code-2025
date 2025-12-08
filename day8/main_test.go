package main

import (
	"bytes"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
	want := 40
	got := PartOne(bytes.Split([]byte(input), []byte("\n")))
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := ""
	want := 1
	got := PartTwo(bytes.Split([]byte(input), []byte("\n")))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}
