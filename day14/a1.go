package day14

import (
	"aoc2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	runAssignment1("day14/input.txt", func(matrix [][]rune) int {
		tilt(matrix, North)
		return calculateLoad(matrix)
	})
}

func runAssignment1(path string, getLoad func(matrix [][]rune) int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	utils.ForEachMatrix(scanner, func(matrix [][]rune) {
		load := getLoad(matrix)
		clipboard.WriteAll(strconv.Itoa(load))
	})
}

func tilt(matrix [][]rune, direction direction) {
	w, h := len(matrix[0]), len(matrix)

	for y := 0; y < w; y++ {
		for x := 0; x < h; x++ {
			switch direction {
			case North:
				if matrix[y][x] == 'O' {
					y_t := y
					for ; y_t > 0; y_t-- {
						if matrix[y_t-1][x] != '.' {
							break
						}
					}
					if y != y_t {
						matrix[y][x] = '.'
						matrix[y_t][x] = 'O'
					}
				}
			case South:
				y_rev := h - y - 1
				if matrix[y_rev][x] == 'O' {
					y_t := y_rev
					for ; y_t < h-1; y_t++ {
						if matrix[y_t+1][x] != '.' {
							break
						}
					}
					if y_rev != y_t {
						matrix[y_rev][x] = '.'
						matrix[y_t][x] = 'O'
					}
				}
			case West:
				if matrix[y][x] == 'O' {
					x_t := x
					for ; x_t > 0; x_t-- {
						if matrix[y][x_t-1] != '.' {
							break
						}
					}
					if x != x_t {
						matrix[y][x] = '.'
						matrix[y][x_t] = 'O'
					}
				}
			case East:
				x_rev := w - x - 1
				if matrix[y][x_rev] == 'O' {
					x_t := x_rev
					for ; x_t < w-1; x_t++ {
						if matrix[y][x_t+1] != '.' {
							break
						}
					}
					if x_rev != x_t {
						matrix[y][x_rev] = '.'
						matrix[y][x_t] = 'O'
					}
				}
			}
		}
	}
}

type direction int

const (
	North direction = 0
	West  direction = 1
	East  direction = 2
	South direction = 3
)

func calculateLoad(matrix [][]rune) int {
	h := len(matrix)
	load := 0
	for y, row := range matrix {
		for _, r := range row {
			if r == 'O' {
				load += h - y
			}
		}
	}
	return load
}
