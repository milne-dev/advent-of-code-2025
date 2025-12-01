package main

import (
	"aoc2025/utils"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	s, err := utils.ReadFileText("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
