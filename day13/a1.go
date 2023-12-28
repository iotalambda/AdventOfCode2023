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

xLoop:
	for x := 0; x < w-1; x++ {
		diff := 0
		for xL, xR := x, x+1; xL >= 0 && xR < w; xL, xR = xL-1, xR+1 {
			for _, row := range pattern {
				if row[xL] != row[xR] {
					diff++
					if diff > smudges {
						continue xLoop
					}
				}
			}
		}
		if diff == smudges {
			return x + 1
		}
	}

yLoop:
	for y := 0; y < h-1; y++ {
		diff := 0
		for yT, yB := y, y+1; yT >= 0 && yB < h; yT, yB = yT-1, yB+1 {
			for x := 0; x < w; x++ {
				if pattern[yT][x] != pattern[yB][x] {
					diff++
					if diff > smudges {
						continue yLoop
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
