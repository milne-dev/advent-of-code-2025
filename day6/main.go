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
	fmt.Println(PartOne(strings.Lines(input)))
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

func PartTwo(fields []string) int {
	var ans int
	return ans
}
