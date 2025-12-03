package main

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s, err := utils.ReadFileText("input.txt")
	if err != nil {
		panic(err)
	}

	fields := strings.Fields(s)
	var banks []int
	for _, field := range fields {
		i, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		banks = append(banks, i)
	}

	fmt.Println(PartOne(banks))
}

func PartOne(banks []int) int {
	fmt.Println(banks)
	return 0
}
