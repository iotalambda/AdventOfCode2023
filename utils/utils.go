package utils

import (
	"fmt"
	"os"
	"strconv"
)

func Ints(source []string) []int {
	result := make([]int, len(source))
	for i, str := range source {
		v, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse int:", err)
			os.Exit(1)
		}
		result[i] = v
	}
	return result
}
