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
					yT := y
					for ; yT > 0; yT-- {
						if matrix[yT-1][x] != '.' {
							break
						}
					}
					if y != yT {
						matrix[y][x] = '.'
						matrix[yT][x] = 'O'
					}
				}
			case South:
				yRev := h - y - 1
				if matrix[yRev][x] == 'O' {
					yT := yRev
					for ; yT < h-1; yT++ {
						if matrix[yT+1][x] != '.' {
							break
						}
					}
					if yRev != yT {
						matrix[yRev][x] = '.'
						matrix[yT][x] = 'O'
					}
				}
			case West:
				if matrix[y][x] == 'O' {
					xT := x
					for ; xT > 0; xT-- {
						if matrix[y][xT-1] != '.' {
							break
						}
					}
					if x != xT {
						matrix[y][x] = '.'
						matrix[y][xT] = 'O'
					}
				}
			case East:
				xRev := w - x - 1
				if matrix[y][xRev] == 'O' {
					xT := xRev
					for ; xT < w-1; xT++ {
						if matrix[y][xT+1] != '.' {
							break
						}
					}
					if xRev != xT {
						matrix[y][xRev] = '.'
						matrix[y][xT] = 'O'
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
