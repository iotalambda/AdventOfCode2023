package utils

import (
	"os"
	"strings"
)

func ReadLines(name string) ([]string, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), Constants.Newline)
	return lines, nil
}
