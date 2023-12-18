package day14

import (
	"fmt"
	"strings"
)

func Assignment2() {
	runAssignment1("day14/input.txt", func(matrix [][]rune) int {
		cycles := 1000000000
		keys := make(map[string]int)
		loads := make([]int, 0)
		for i := 0; ; i++ {
			tilt(matrix, North)
			tilt(matrix, West)
			tilt(matrix, South)
			tilt(matrix, East)
			key := key(matrix)
			if start, found := keys[key]; found {
				return loads[((cycles-start-1)%(i-start))+start]
			} else {
				keys[key] = i
				loads = append(loads, calculateLoad(matrix))
			}

		}
	})
}

func key(matrix [][]rune) string {
	rocks := make([]string, 0)
	for y, row := range matrix {
		for x, r := range row {
			if r == 'O' {
				rocks = append(rocks, fmt.Sprintf("%d,%d", y, x))
			}
		}
	}
	return strings.Join(rocks, ";")
}
