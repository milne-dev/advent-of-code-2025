package main

import (
	"aoc2025/utils"
	"bytes"
	"fmt"
)

func main() {
	input := utils.ReadInput()
	fmt.Println(PartOne(bytes.Split([]byte(input), []byte("\n"))))
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

func PartTwo(lines []string) int {
	return 0
}
