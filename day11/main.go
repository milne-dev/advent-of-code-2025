package main

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadInput()
	lines := utils.StringLines(input)
	fmt.Println(PartTwo(lines))
}

func PartOne(lines []string) int {
	adj := make(map[string][]string)
	for _, line := range lines {
		adj[line[0:3]] = strings.Fields(line[5:])
	}
	return search(adj, adj["you"])
}

func search(adj map[string][]string, edges []string) int {
	var ans int
	for _, edge := range edges {
		if edge == "out" {
			return 1
		}
		ans += search(adj, adj[edge])
	}
	return ans
}

const adjSize = 17576 // "ZZZ" + 1

type rec struct {
	edge     uint16
	dac, fft bool
}

func PartTwo(lines []string) int {
	adj := make([][]uint16, adjSize)
	for _, line := range lines {
		fields := strings.Fields(line[5:])
		k := TripletToInt(line[0:3])
		adj[k] = make([]uint16, len(fields))
		for i, field := range fields {
			adj[k][i] = TripletToInt(field)
		}
	}
	memo := make(map[rec]int)
	return searchP2(adj, adj[TripletToInt("svr")], false, false, memo)
}

var (
	OUT = TripletToInt("out")
	DAC = TripletToInt("dac")
	FFT = TripletToInt("fft")
)

func searchP2(adj [][]uint16, edges []uint16, dac, fft bool, memo map[rec]int) int {
	var ans int
	for _, edge := range edges {
		if edge == OUT {
			if dac && fft {
				return 1
			}
			return 0
		}

		if v, ok := memo[rec{edge, dac, fft}]; ok {
			return v
		}

		var edgeVal int
		if edge == DAC {
			edgeVal = searchP2(adj, adj[edge], true, fft, memo)
		} else if edge == FFT {
			edgeVal = searchP2(adj, adj[edge], dac, true, memo)
		} else {
			edgeVal = searchP2(adj, adj[edge], dac, fft, memo)
		}

		ans += edgeVal
		memo[rec{edge, dac, fft}] = edgeVal
	}
	return ans
}

func TripletToInt(s string) uint16 {
	var n uint16
	for i := range s {
		n = n*26 + uint16(s[i]-'a')
	}
	return n
}
