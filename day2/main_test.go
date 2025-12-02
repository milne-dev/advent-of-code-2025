package main

import "testing"

func TestPartOne(t *testing.T) {
	d := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
		"565653-565659",
		"824824821-824824827",
		"2121212118-2121212124",
	}
	want := 1227775554
	got := PartOne(d)
	if want != got {
		t.Errorf(`PartOne() want %v, got %v`, want, got)
	}
}

func TestIsInvalid(t *testing.T) {
	d := []struct {
		id  int
		res bool
	}{
		{22, true},
		{21, false},
		{101, false},
		{101010, false},
		{101101, true},
		{55555, false},
		{555555, true},
		{5555555, false},
		{55655, false},
		{555655, false},
		{5556555, false},
		{6555555, false},
		{7, false},
		{721, false},
		{727, false},
		{7227, false},
		{7272, true},
	}

	for _, c := range d {
		want := c.res
		got := IsInvalid(c.id)
		if want != got {
			t.Errorf(`IsInvalid(%v) want %v, got %v`, c.id,
				want, got)
		}
	}
}

func TestPartTwo(t *testing.T) {
	d := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
		"565653-565659",
		"824824821-824824827",
		"2121212118-2121212124",
	}
	want := 4174379265
	got := PartTwo(d)
	if want != got {
		t.Errorf(`PartTwo() want %v, got %v`, want, got)
	}
}

func TestIsInvalidP2(t *testing.T) {
	d := []struct {
		id  int
		res bool
	}{
		{22, true},
		{21, false},
		{101, false},
		{101010, true},
		{101101, true},
		{55555, true},
		{555555, true},
		{5555555, true},
		{55655, false},
		{555655, false},
		{5556555, false},
		{6555555, false},
		{7, false},
		{777, true},
		{721, false},
		{727, false},
		{7227, false},
		{7272, true},
		{123123, true},
		{123123123, true},
		{1231231234, false},
		{1234123123, false},
	}

	for _, c := range d {
		want := c.res
		got := IsInvalidP2(c.id)
		if want != got {
			t.Errorf(`IsInvalidP2(%v) want %v, got %v`, c.id,
				want, got)
		}
	}
}
