package main

import (
	"aoc2025/utils"
	"bytes"
	"fmt"
)

func main() {
	input := utils.ReadInput()
	fmt.Println(PartTwo(bytes.Split([]byte(input), []byte("\n"))))
}

func PartOne(lines [][]byte) int {
	var ans int
	var q [][2]int

	for i, line := range lines {
		for j, c := range line {
			if c == 'S' {
				q = [][2]int{{i + 1, j}}
				break
			}
		}
	}

	for len(q) > 0 {
		var nq [][2]int

		for _, pos := range q {
			i, j := pos[0], pos[1]

			if i == len(lines) || j == len(lines[0]) {
				continue
			}

			switch lines[i][j] {
			case '|':
				continue
			case '^':
				ans++
				nq = append(nq, [2]int{i, j - 1})
				nq = append(nq, [2]int{i, j + 1})
			case '.':
				lines[i][j] = '|'
				nq = append(nq, [2]int{i + 1, j})
			}
		}

		q = nq
	}

	return ans
}

type rec struct {
	i, j int
	path string
}

func PartTwo(lines [][]byte) int {
	dp := make([][]int, len(lines)+1)
	for i := range dp {
		dp[i] = make([]int, len(lines[0])+1)
	}

	for i := len(lines) - 1; i >= 0; i-- {
		for j := range lines[0] {
			if lines[i][j] != '^' {
				dp[i][j] = dp[i+1][j]
			}
		}

		for j := range lines[0] {
			if lines[i][j] != '^' {
				continue
			}

			var left, right int

			if j > 0 {
				left = dp[i][j-1]
			}

			if j < len(lines[0])-1 {
				right = dp[i][j+1]
			}

			dp[i][j] = 1 + left + right
		}
	}

	for i, line := range lines {
		for j, c := range line {
			if c == 'S' {
				return 1 + dp[i][j]
			}
		}
	}

	return 0
}
