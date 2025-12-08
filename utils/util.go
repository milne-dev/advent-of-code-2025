package utils

import (
	"bytes"
	"os"
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

func ByteLines(input []byte) [][]byte {
	return bytes.Split(input, []byte("\n"))
}
