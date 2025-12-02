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

	fields := strings.Split(s, ",")

	fmt.Println(PartOne(fields))
}

func PartOne(fields []string) int {
	var ans int
	for _, field := range fields {
		parts := strings.Split(field, "-")
		if IsInvalid(parts[0]) {
			n, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			ans += n
		}
		if IsInvalid(parts[1]) {
			n, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			ans += n
		}
	}
	return ans
}

func IsInvalid(id string) bool {
	if len(id)&1 == 1 {
		return false
	}

	return true
}
