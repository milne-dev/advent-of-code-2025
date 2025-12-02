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

	fields := strings.Split(s, ",")

	fmt.Println(PartOne(fields))
}

func PartOne(fields []string) int {
	return 0
}
