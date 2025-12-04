package main

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(PartTwoStack(Setup()))
}

func Setup() []string {
	return strings.Fields(utils.ReadInput())
}

func PartTwoStack(banks []string) int {
	var ans int
	for _, bank := range banks {
		ans += MaxNumTwelveStack(bank)
	}
	return ans
}

func MaxNumTwelveStack(bank string) int {
	stack := make([]int, 1, 12)
	stack[0] = int(bank[0] - '0')

	for i := 1; i < len(bank); i++ {
		d := int(bank[i] - '0')
		shouldAppend := len(stack) < 12

		for len(stack) > 0 && d > stack[len(stack)-1] && len(stack)+len(bank)-i > 12 {
			shouldAppend = true
			stack = stack[:len(stack)-1]
		}

		if shouldAppend {
			stack = append(stack, d)
		}
	}

	var ans int
	for _, n := range stack {
		ans = ans*10 + n
	}
	return ans
}

func PartTwo(banks []string) int {
	var ans int
	for _, bank := range banks {
		ans += MaxNumTwelve(bank)
	}
	return ans
}

func MaxNumTwelve(bank string) int {
	lookup := [10][]int{}
	for i, d := range bank {
		lookup[d-'0'] = append(lookup[d-'0'], i)
	}

	for i := 9; i >= 0; i-- {
		for _, j := range lookup[i] {
			res := search(j, 1, i, lookup)
			if res != 0 {
				return res
			}
		}
	}

	return 0
}

func search(k, size, cur int, lookup [10][]int) int {
	if size == 12 {
		return cur
	}

	for i := 9; i >= 0; i-- {
		for _, j := range lookup[i] {
			if k >= j {
				continue
			}
			res := search(j, size+1, cur*10+i, lookup)
			if res != 0 {
				return res
			}
		}
	}

	return 0
}

func PartOne(banks []string) int {
	var ans int
	for _, bank := range banks {
		ans += MaxNum(bank)
	}
	return ans
}

func MaxNum(bank string) int {
	prefix := make([]int, len(bank))
	suffix := make([]int, len(bank))

	prefix[0] = int(bank[0] - '0')
	for i := 1; i < len(bank); i++ {
		prefix[i] = max(prefix[i-1], int(bank[i]-'0'))
	}

	suffix[len(suffix)-1] = int(bank[len(bank)-1] - '0')
	for i := len(bank) - 2; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], int(bank[i]-'0'))
	}

	var ans int
	for i := 0; i < len(bank)-1; i++ {
		ans = max(ans, prefix[i]*10+suffix[i+1])
	}
	return ans
}
