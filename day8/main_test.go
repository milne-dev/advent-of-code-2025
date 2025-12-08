package main

import (
	"bytes"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := ""
	want := 1
	got := PartOne(bytes.Split([]byte(input), []byte("\n")))
	if want != got {
		t.Errorf(`PartOne(...) want %v, got %v`,
			want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input := ""
	want := 1
	got := PartTwo(bytes.Split([]byte(input), []byte("\n")))
	if want != got {
		t.Errorf(`PartTwo(...) want %v, got %v`,
			want, got)
	}
}
