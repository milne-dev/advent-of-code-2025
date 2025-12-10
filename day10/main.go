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
	var ans int

	for i, line := range lines {
		fmt.Printf("line %d of %d\n", i+1, len(lines))
		ans += processLineP2(line)
	}

	return ans
}

func processLineP2(line string) int {
	var target [10]uint16
	var ti int
	var buttons []Button

	mode := ModeTarget
	var buttonGroup Button
	var buttonNum int

	var voltageNum uint16

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

		case ModeButtons:
			if c == '{' {
				mode = ModeVoltage
				continue
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

		case ModeVoltage:
			if c == '}' {
				target[ti] = voltageNum
				ti++
				continue
			}

			if c == ',' {
				target[ti] = voltageNum
				ti++
				voltageNum = 0
			} else {
				voltageNum = voltageNum*10 + uint16(c-'0')
			}
		}
	}

	q := make(map[[10]uint16]int)
	q[[10]uint16{}] = 0
	var passes int
	targetSum := sumTarget(target)

	for len(q) > 0 {
		nq := make(map[[10]uint16]int)

		for current, sum := range q {
			if sum > targetSum {
				continue
			}

			if current == target {
				return passes
			}

			// send this one back to the q with each possible btn press
			for _, b := range buttons {
				clone := current
				applyBtnToVoltage(&clone, b)
				nq[clone] = sum + len(b)
			}
		}

		q = nq
		passes++
	}

	return 0
}

func sumTarget(a [10]uint16) int {
	var n int
	for _, x := range a {
		n += int(x)
	}
	return n
}

func applyBtnToVoltage(voltage *[10]uint16, button Button) {
	for _, i := range button {
		voltage[i]++
	}
}
