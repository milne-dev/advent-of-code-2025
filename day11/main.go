package main

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadInput()
	lines := utils.StringLines(input)
	fmt.Println(PartOne(lines))
}

func PartOne(lines []string) int {
	adj := make(map[string][]string)
	for _, line := range lines {
		adj[line[0:4]] = strings.Fields(line[5:])
	}
	return 0
}
