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
	paths := make(map[string]bool)
	var q []rec

	for i, line := range lines {
		for j, c := range line {
			if c == 'S' {
				q = []rec{
					{i: i + 1, j: j, path: coordsToS(i, j)},
				}
				break
			}
		}
	}

	for len(q) > 0 {
		var nq []rec

		for _, r := range q {
			i, j := r.i, r.j

			if i == len(lines) {
				paths[r.path] = true
				continue
			}

			if j == len(lines[0]) {
				continue
			}

			switch lines[i][j] {
			case '|':
				continue
			case '^':
				nq = append(nq, rec{i: i, j: j - 1, path: r.path + coordsToS(i, j)})
				nq = append(nq, rec{i: i, j: j + 1, path: r.path + coordsToS(i, j)})
			case '.':
				nq = append(nq, rec{i: i + 1, j: j, path: r.path + coordsToS(i, j)})
			}
		}

		q = nq
	}

	return len(paths)
}

func coordsToS(x, y int) string {
	return fmt.Sprintf("%v,%v,", x, y)
}
