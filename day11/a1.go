package day11

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	runAssignment1("day11/input.txt", 2)
}

func runAssignment1(path string, expansion_rate int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	galaxies := make([][2]int, 0)
	y_max, x_max := 0, 0
	for y := 0; scanner.Scan(); y++ {
		for x, r := range scanner.Text() {
			if r == '#' {
				galaxies = append(galaxies, [...]int{y, x})
				y_max = max(y_max, y)
				x_max = max(x_max, x)
			}
		}
	}

	y_galaxies_lookup := make(map[int]struct{}, 0)
	x_galaxies_lookup := make(map[int]struct{}, 0)
	for _, galaxy := range galaxies {
		y_galaxies_lookup[galaxy[0]] = struct{}{}
		x_galaxies_lookup[galaxy[1]] = struct{}{}
	}

	y_voids := make([]int, 0)
y_max_loop:
	for y := 0; y < y_max; y++ {
		for y_galaxy := range y_galaxies_lookup {
			if y == y_galaxy {
				continue y_max_loop
			}
		}
		y_voids = append(y_voids, y)
	}
	slices.SortFunc(y_voids, func(a, b int) int {
		return b - a
	})

	x_voids := make([]int, 0)
x_max_loop:
	for x := 0; x < x_max; x++ {
		for x_galaxy := range x_galaxies_lookup {
			if x == x_galaxy {
				continue x_max_loop
			}
		}
		x_voids = append(x_voids, x)
	}
	slices.SortFunc(x_voids, func(a, b int) int {
		return b - a
	})

	for _, y_void := range y_voids {
		for g_ix, galaxy := range galaxies {
			if y_curr := galaxy[0]; y_curr > y_void {
				galaxies[g_ix][0] = y_curr + expansion_rate - 1
			}
		}
	}
	for _, x_void := range x_voids {
		for g_ix, galaxy := range galaxies {
			if x_curr := galaxy[1]; x_curr > x_void {
				galaxies[g_ix][1] = x_curr + expansion_rate - 1
			}
		}
	}

	distances_total := 0
	for g1_ix, g1 := range galaxies {
		for _, g2 := range galaxies[g1_ix+1:] {
			d_y := (int)(math.Abs((float64)(g2[0] - g1[0])))
			d_x := (int)(math.Abs((float64)(g2[1] - g1[1])))
			distances_total += d_y + d_x
		}
	}

	clipboard.WriteAll(strconv.Itoa(distances_total))
}
