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
	memo := make(map[[2]int]int)
	for i, line := range lines {
		for j, c := range line {
			if c == 'S' {
				return 1 + search(lines, i+1, j, memo)
			}
		}
	}

	return 0
}

func search(lines [][]byte, i, j int, memo map[[2]int]int) (res int) {
	if i == len(lines) || j == len(lines[0]) {
		return 0
	}

	if v, ok := memo[[2]int{i, j}]; ok {
		return v
	}

	if lines[i][j] == '^' {
		res = 1 + search(lines, i, j-1, memo) + search(lines, i, j+1, memo)
	} else {
		res = search(lines, i+1, j, memo)
	}

	memo[[2]int{i, j}] = res

	return
}

func coordsToS(x, y int) string {
	return fmt.Sprintf("%v,%v,", x, y)
}
