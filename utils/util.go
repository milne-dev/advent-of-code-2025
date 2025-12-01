package utils

import "os"

func ReadFileText(path string) (string, error) {
	b, err := os.ReadFile(path)
	return string(b), err
}
