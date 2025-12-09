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

	// lets add test cases from the web...
	input = `1,0
3,0
3,6
16,6
16,0
18,0
18,9
13,9
13,7
6,7
6,9
1,9`
	want = 30
	got = PartTwo(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}

	input = `1,1
8,1
8,3
3,3
3,4
8,4
8,9
18,9
18,11
5,11
5,9
4,9
4,11
1,11
1,7
6,7
6,6
1,6`
	want = 88
	got = PartTwo(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}

	input = `1,5
3,5
3,8
7,8
7,5
9,5
9,10
11,10
11,3
6,3
6,7
4,7
4,1
13,1
13,12
1,12`
	want = 72
	got = PartTwo(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}
