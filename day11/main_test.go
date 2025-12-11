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

// string-lookup BenchmarkPartTwoSearch-8   	 3406622	       350.8 ns/op
// uint16-lookup BenchmarkPartTwoSearch-8   	16830844	        69.40 ns/op
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
	adj := make([][]uint16, adjSize)
	for _, line := range lines {
		fields := strings.Fields(line[5:])
		k := TripletToInt(line[0:3])
		adj[k] = make([]uint16, len(fields))
		for i, field := range fields {
			adj[k][i] = TripletToInt(field)
		}
	}

	svr := TripletToInt("svr")

	for b.Loop() {
		searchP2(adj, adj[svr], false, false)
	}
}

func TestTripletToInt(t *testing.T) {
	if TripletToInt("AAF") == TripletToInt("FAA") || TripletToInt("AAF") == TripletToInt("AFA") {
		t.Errorf("bad TestTripletToInt")
	}
	if TripletToInt("BAA") == TripletToInt("AKA") {
		t.Errorf("I thought something was off")
	}
}
