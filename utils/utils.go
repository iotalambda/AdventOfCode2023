package utils

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func ForEachMatrix(scanner *bufio.Scanner, action func(matrix [][]rune)) {
	var matrix [][]rune
	for scanner.Scan() {
		row := []rune(scanner.Text())

		if len(row) <= 1 {
			action(matrix)
			matrix = nil
			continue
		}

		if matrix == nil {
			matrix = make([][]rune, 0)
		}

		matrix = append(matrix, row)
	}
	action(matrix)
}

func Prepend[T any](slice []T, elems ...T) []T {
	slices.Reverse(slice)
	slice = append(slice, elems...)
	slices.Reverse(slice)
	return slice
}
