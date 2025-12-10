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

const (
	ModeTarget = iota
	ModeButtons
	ModeVoltage
)

func PartOne(lines []string) int {
	var ans int

	for _, line := range lines {
		var target []bool
		var buttons [][]int

		mode := ModeTarget
		var buttonGroup []int
		var buttonNum int

		for i := 1; i < len(line); i++ {
			c := line[i]

			if c == ' ' || c == '(' {
				continue
			}

			switch mode {
			case ModeTarget:
				if c == ']' {
					mode = ModeButtons
					continue
				}

				if c == '.' {
					target = append(target, false)
				} else {
					target = append(target, true)
				}
			case ModeButtons:
				// don't handle voltage for this problem
				if c == '{' {
					break
				}

				switch c {
				case ',':
					buttonGroup = append(buttonGroup, buttonNum)
					buttonNum = 0
				case ')':
					buttonGroup = append(buttonGroup, buttonNum)
					buttonNum = 0
					buttons = append(buttons, buttonGroup)
					buttonGroup = []int{}
				default:
					buttonNum = buttonNum*10 + int(c-'0')
				}
			}
		}

		fmt.Println(target)
		fmt.Println(buttons)
	}

	return ans
}
