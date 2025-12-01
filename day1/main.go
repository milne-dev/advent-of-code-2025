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

	lines := strings.Fields(s)
	fmt.Println(Solve(lines))
}

func Solve(input []string) int {
	pos := 50
	var ans int

	for _, l := range input {
		var n int
		for i := 1; i < len(l); i++ {
			n *= 10
			n += int(l[i] - '0')
		}

		if l[0] == 'L' {
			pos = ChangePos(pos, -n)
		} else {
			pos = ChangePos(pos, n)
		}

		if pos == 0 {
			ans++
		}
	}

	return ans
}

func ChangePos(p, n int) int {
	if n >= 0 {
		return (p + n) % 100
	}

	n = (-n) % 100
	x := p - n
	if x >= 0 {
		return x
	}
	return 100 + x
}
