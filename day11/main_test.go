package main

import (
	"aoc2025/utils"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`
	want := 5
	got := PartOne(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
	want := 2
	got := PartTwo(utils.StringLines(input))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}

// string-based BenchmarkPartTwoSearch-8   	 3406622	       350.8 ns/op
func BenchmarkPartTwoSearch(b *testing.B) {
	input := `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
	lines := utils.StringLines(input)
	adj := make(map[string][]string)
	for _, line := range lines {
		adj[line[0:3]] = strings.Fields(line[5:])
	}

	for b.Loop() {
		searchP2(adj, adj["svr"], false, false)
	}
}
