package main

import (
	"aoc2025/utils"
	"cmp"
	"fmt"
	"math"
	"slices"
)

func main() {
	input := utils.ReadInput()
	lines := utils.StringLines(input)
	fmt.Println(PartOne(lines))
}

type pos struct {
	x, y, z int
}

type dist struct {
	v    float64
	p, q pos
}

func PartOne(lines []string) int {
	positions := make([]pos, len(lines))

	for i, line := range lines {
		var p pos
		_, err := fmt.Sscanf(line, "%d,%d,%d", &p.x, &p.y, &p.z)
		if err != nil {
			panic(err)
		}
		positions[i] = p
	}

	var distances []dist

	for i, p := range positions {
		for j := i + 1; j < len(positions); j++ {
			q := positions[j]
			distances = append(distances, dist{v: distance(p, q), p: p, q: q})
		}
	}

	slices.SortFunc(distances, func(a, b dist) int {
		return cmp.Compare(a.v, b.v)
	})

	fmt.Println(distances[0])

	return 0
}

func distance(p, q pos) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2) + math.Pow(float64(p.z-q.z), 2))
}

func PartTwo(lines [][]byte) int {
	return 0
}
