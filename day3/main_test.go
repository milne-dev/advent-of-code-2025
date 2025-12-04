package main

import (
	"aoc2025/utils"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	want := 357
	got := PartOne(input)
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	want := 3121910778619
	got := PartTwo(input)
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}

// BenchmarkPartTwo-8   	     904	   1318104 ns/op
func BenchmarkPartTwo(b *testing.B) {
	s, err := utils.ReadFileText("input.txt")
	if err != nil {
		panic(err)
	}

	banks := strings.Fields(s)

	for b.Loop() {
		PartTwo(banks)
	}
}
