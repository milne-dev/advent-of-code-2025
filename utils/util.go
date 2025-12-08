package utils

import (
	"bytes"
	"os"
	"strings"
)

func ReadInput() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func ReadInputBytes() []byte {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return b
}

func StringLines(input string) []string {
	return strings.Split(input, "\n")
}

func ByteLines(input []byte) [][]byte {
	return bytes.Split(input, []byte("\n"))
}
