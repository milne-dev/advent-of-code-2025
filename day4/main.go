package main

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func main() {
	s, err := utils.ReadFileText("input.txt")
	if err != nil {
		panic(err)
	}

	grid := strings.Fields(s)
	fmt.Println(PartOne(grid))
}

func PartOne(grid []string) int {
	return 0
}
