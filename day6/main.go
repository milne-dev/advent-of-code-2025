package main

import (
	"aoc2025/utils"
	"fmt"
	"iter"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadInput()
	fmt.Println(PartTwo(strings.Split(input, "\n")))
}

func PartOne(lines iter.Seq[string]) int {
	var ans int

	nums := make([][]int, 0)
	var cmds []byte

	for line := range lines {
		lineFields := strings.Fields(line)

		if lineFields[0][0] == '*' || lineFields[0][0] == '+' {
			cmds = make([]byte, len(lineFields))
			for i, field := range lineFields {
				cmds[i] = field[0]
			}
			break
		}

		numsLine := []int{}
		for _, field := range lineFields {
			n, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			numsLine = append(numsLine, n)
		}
		nums = append(nums, numsLine)
	}

	for i, cmd := range cmds {
		var x int
		if cmd == '*' {
			x = nums[0][i]
			for j := 1; j < len(nums); j++ {
				x *= nums[j][i]
			}
		} else {
			for j := range nums {
				x += nums[j][i]
			}
		}
		ans += x
	}

	return ans
}

func PartTwo(lines []string) int {
	var ans int

	var nums []int
	var n int

	for j := len(lines[0]) - 1; j >= 0; j-- {
		for i := range lines {
			c := lines[i][j]
			if c == '*' || c == '+' {
				nums = append(nums, n)
				n = 0

				var x int
				if c == '*' {
					x = nums[0]
					for k := 1; k < len(nums); k++ {
						x *= nums[k]
					}
				} else {
					for _, n := range nums {
						x += n
					}
				}
				ans += x
				nums = []int{}
				j--
				continue
			}

			if c >= '0' && c <= '9' {
				n = n*10 + int(c-'0')
			}

			if i == len(lines)-1 {
				nums = append(nums, n)
				n = 0
			}
		}
	}

	return ans
}
