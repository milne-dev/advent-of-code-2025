package main

import (
	"aoc2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput()
	lines := utils.StringLines(input)
	fmt.Println(PartTwo(lines))
}

const (
	ModeTarget = iota
	ModeButtons
	ModeVoltage
)

type qrec struct {
	current []bool
	button  []int
}

type Button []int

func PartOne(lines []string) int {
	var ans int

	for _, line := range lines {
		ans += processLine(line)
	}

	return ans
}

func processLine(line string) int {
	var targetArr []bool
	var target int
	var buttons []Button

	mode := ModeTarget
	var buttonGroup Button
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
				targetArr = append(targetArr, false)
			} else {
				targetArr = append(targetArr, true)
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
				buttonGroup = Button{}
			default:
				buttonNum = buttonNum*10 + int(c-'0')
			}
		}
	}

	for i := len(targetArr) - 1; i >= 0; i-- {
		target <<= 1
		if targetArr[i] {
			target |= 1
		}
	}

	var passes int
	q := make(map[int]struct{})

	for _, button := range buttons {
		q[applyButton(0, button)] = struct{}{}
	}

outer:
	for len(q) > 0 {
		nq := make(map[int]struct{})
		passes++

		for current := range q {
			if current == target {
				break outer
			}

			// send this one back to the q with each possible btn press
			for _, b := range buttons {
				nq[applyButton(current, b)] = struct{}{}
			}
		}

		q = nq
	}

	return passes
}

func applyButton(current int, b Button) int {
	for _, i := range b {
		mask := 1 << i
		current ^= mask
	}
	return current
}

func PartTwo(lines []string) int {
	return 0
}
