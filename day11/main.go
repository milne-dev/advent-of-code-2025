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
	return 0
}
