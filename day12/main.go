package main

import (
	"aoc2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput()
	lines := utils.StringLines(input)
	fmt.Println(PartOne(lines))
}

func PartOne(lines []string) int {
	return 0
}

type Present [3][3]bool

func RotatePresent(p Present) {

}

func boolToSym(b bool) string {
	if b {
		return "#"
	}
	return "."
}
