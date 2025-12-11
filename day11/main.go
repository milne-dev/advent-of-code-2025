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

func PartTwo(lines []string) int {
	adj := make(map[string][]string)
	for _, line := range lines {
		adj[line[0:3]] = strings.Fields(line[5:])
	}
	return searchP2(adj, adj["svr"], false, false)
}

func searchP2(adj map[string][]string, edges []string, dac, fft bool) int {
	var ans int
	for _, edge := range edges {
		if edge == "out" {
			if dac && fft {
				return 1
			}
			return 0
		}
		if edge == "dac" {
			ans += searchP2(adj, adj[edge], true, fft)
		} else if edge == "fft" {
			ans += searchP2(adj, adj[edge], dac, true)
		} else {
			ans += searchP2(adj, adj[edge], dac, fft)
		}
	}
	return ans
}
