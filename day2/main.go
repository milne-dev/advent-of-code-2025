package main

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fields := strings.Split(utils.ReadInput(), ",")
	fmt.Println(PartTwo(fields))
}

func PartOne(fields []string) int {
	var ans int
	for _, field := range fields {
		parts := strings.Split(field, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for ; start <= end; start++ {
			if IsInvalid(start) {
				ans += start
			}
		}
	}
	return ans
}

func IsInvalid(idInt int) bool {
	id := strconv.Itoa(idInt)
	if len(id)&1 == 1 {
		return false
	}
	return id[0:len(id)/2] == id[len(id)/2:]
}

func PartTwo(fields []string) int {
	var ans int
	for _, field := range fields {
		parts := strings.Split(field, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for ; start <= end; start++ {
			if IsInvalidP2(start) {
				ans += start
			}
		}
	}
	return ans
}

func IsInvalidP2(idInt int) bool {
	id := strconv.Itoa(idInt)
	for j := 1; j <= len(id)/2; j++ {
		if isRepeatingForEntirety(id, id[0:j], j) {
			return true
		}
	}
	return false
}

func isRepeatingForEntirety(id, sub string, inc int) bool {
	if len(id)%inc != 0 {
		return false
	}
	for i := inc; i+inc <= len(id); i += inc {
		if sub != id[i:i+inc] {
			return false
		}
	}

	return true
}
