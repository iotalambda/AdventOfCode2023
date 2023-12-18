package day13

import (
	"aoc2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	runAssignment1("day13/input.txt", 0)
}

func runAssignment1(path string, smudges int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0
	utils.ForEachMatrix(scanner, func(matrix [][]rune) {
		total += summarize(matrix, smudges)
	})

	clipboard.WriteAll(strconv.Itoa(total))
}

func summarize(pattern [][]rune, smudges int) int {
	w, h := len(pattern[0]), len(pattern)

x_loop:
	for x := 0; x < w-1; x++ {
		diff := 0
		for x_l, x_r := x, x+1; x_l >= 0 && x_r < w; x_l, x_r = x_l-1, x_r+1 {
			for _, row := range pattern {
				if row[x_l] != row[x_r] {
					diff++
					if diff > smudges {
						continue x_loop
					}
				}
			}
		}
		if diff == smudges {
			return x + 1
		}
	}

y_loop:
	for y := 0; y < h-1; y++ {
		diff := 0
		for y_t, y_b := y, y+1; y_t >= 0 && y_b < h; y_t, y_b = y_t-1, y_b+1 {
			for x := 0; x < w; x++ {
				if pattern[y_t][x] != pattern[y_b][x] {
					diff++
					if diff > smudges {
						continue y_loop
					}
				}
			}
		}
		if diff == smudges {
			return 100 * (y + 1)
		}
	}

	return 0
}
