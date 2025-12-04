package main

import (
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
	banks := Setup()
	for b.Loop() {
		PartTwo(banks)
	}
}

func TestPartTwoStack(t *testing.T) {
	input := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}
	want := 3121910778619
	got := PartTwoStack(input)
	if want != got {
		t.Errorf(`PartTwoStack(...) want %v, got %v`,
			want, got)
	}
}

// BenchmarkPartTwoStack-8   	   22904	     50667 ns/op
func BenchmarkPartTwoStack(b *testing.B) {
	banks := Setup()
	for b.Loop() {
		PartTwoStack(banks)
	}
}
