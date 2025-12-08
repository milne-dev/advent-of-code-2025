package main

import (
	"aoc2025/utils"
	"cmp"
	"fmt"
	"maps"
	"math"
	"slices"
)

func main() {
	input := utils.ReadInput()
	lines := utils.StringLines(input)
	fmt.Println(PartTwo(lines))
}

type pos struct {
	x, y, z int
}

type dist struct {
	v    float64
	p, q pos
}

func PartOne(lines []string, n int) int {
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
	
	var conns []map[pos]bool

	for i := range n {
		d := distances[i]
		p, q := d.p, d.q

		s1, s2 := -1, -1

		for i, con := range conns {
			if con[p] {
				s1 = i
			} else if con[q] {
				s2 = i
			}
		}

		if s1 == -1 && s2 == -1 {
			m := make(map[pos]bool)
			m[p] = true
			m[q] = true
			conns = append(conns, m)
		} else if s1 == s2 {
			continue
		} else if s1 != -1 && s2 != -1 {
			// merge sets
			maps.Copy(conns[s1], conns[s2])
			conns = append(conns[:s2], conns[s2+1:]...)
		} else if s1 != -1 {
			conns[s1][q] = true
		} else {
			conns[s2][p] = true
		}
	}

	sizes := make([]int, 0, len(conns))
	for _, v := range conns {
		sizes = append(sizes, len(v))
	}

	slices.Sort(sizes)

	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func distance(p, q pos) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2) + math.Pow(float64(p.z-q.z), 2))
}

func PartTwo(lines []string) int {
	return 0
}
