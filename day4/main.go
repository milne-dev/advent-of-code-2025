package main

import (
	"aoc2025/utils"
	"fmt"
)

func main() {
	s, err := utils.ReadFileText("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(PartOne(s))
}

func PartOne(input string) int {
	return 0
}
