package main

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func main() {
	lines := strings.Fields(utils.ReadInput())
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
			ans += TimesPastZero(pos, -n)
			pos = ChangePos(pos, -n)
		} else {
			ans += TimesPastZero(pos, n)
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

// couldn't find a clever way to not simulate this
func TimesPastZero(p, n int) int {
	var ans int
	if n > 0 {
		for ; n > 1; n-- {
			p++
			if p%100 == 0 {
				ans++
			}
		}
	} else {
		for ; n < -1; n++ {
			p--
			if p%100 == 0 {
				ans++
			}
		}
	}
	return ans
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
