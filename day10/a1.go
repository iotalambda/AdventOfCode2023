package day10

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment1() {

	area := getArea("day10/input.txt")
	fp := '#'
	finders := getInitialFinders(area, 'S')
	step := buildStep(area, &finders, fp)
	distance := runAssignment1(area, &finders, fp, step)
	clipboard.WriteAll(strconv.Itoa(distance))
}

func runAssignment1(area [][]rune, finders *[]finder, footprint rune, step func(f_ix int, x int, y int, look bool)) int {
	canGoUp := func(f finder, accept_start bool) bool {
		return pointsUp(f.r, true) && pointsDown(area[f.y-1][f.x], accept_start)
	}
	canGoDown := func(f finder, accept_start bool) bool {
		return pointsDown(f.r, true) && pointsUp(area[f.y+1][f.x], accept_start)
	}
	canGoLeft := func(f finder, accept_start bool) bool {
		return pointsLeft(f.r, true) && pointsRight(area[f.y][f.x-1], accept_start)
	}
	canGoRight := func(f finder, accept_start bool) bool {
		return pointsRight(f.r, true) && pointsLeft(area[f.y][f.x+1], accept_start)
	}

	return runFinders(area, finders, footprint, step, canGoUp, canGoDown, canGoLeft, canGoRight)
}

func pointsUp(r rune, or_start bool) bool {
	return r == 'J' || r == '|' || r == 'L' || or_start && r == 'S'
}
func pointsDown(r rune, or_start bool) bool {
	return r == '7' || r == '|' || r == 'F' || or_start && r == 'S'
}
func pointsLeft(r rune, or_start bool) bool {
	return r == 'J' || r == '-' || r == '7' || or_start && r == 'S'
}
func pointsRight(r rune, or_start bool) bool {
	return r == 'L' || r == '-' || r == 'F' || or_start && r == 'S'
}

type finder struct {
	x int
	y int
	r rune
}

func getArea(file_path string) [][]rune {
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	area := make([][]rune, 0)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		area = append(area, line)
	}

	return area
}

func getInitialFinders(area [][]rune, start_at rune) []finder {
	s_x, s_y := -1, -1
	for _, line := range area {
		if s_x == -1 {
			s_x = slices.Index(line, start_at)
			s_y++
		}
	}
	return []finder{{s_x, s_y, start_at}}
}

func buildStep(area [][]rune, finders *[]finder, footprint rune) func(f_ix int, x int, y int, look bool) {
	step := func(f_ix int, x int, y int, look bool) {
		next := finder{x, y, area[y][x]}
		if look {
			(*finders)[f_ix] = next
		} else {
			*finders = append(*finders, next)
		}
		area[y][x] = footprint
	}
	return step
}

func runFinders(area [][]rune, finders *[]finder, footprint rune, step func(f_ix int, x int, y int, look bool),
	canGoUp func(finder, bool) bool,
	canGoDown func(finder, bool) bool,
	canGoLeft func(finder, bool) bool,
	canGoRight func(finder, bool) bool) int {

	w_area := len(area[0])
	h_area := len(area)

	var d int

	for ; len(*finders) > 0; d++ {
		var deleted_finders []int
		for f_ix, f := range *finders {
			look := true
			if f.y > 0 && canGoUp(f, d > 1) {
				step(f_ix, f.x, f.y-1, look)
				look = false
			}
			if f.y < h_area-1 && canGoDown(f, d > 1) {
				step(f_ix, f.x, f.y+1, look)
				look = false
			}
			if f.x > 0 && canGoLeft(f, d > 1) {
				step(f_ix, f.x-1, f.y, look)
				look = false
			}
			if f.x < w_area-1 && canGoRight(f, d > 1) {
				step(f_ix, f.x+1, f.y, look)
				look = false
			}

			if look {
				if deleted_finders == nil {
					deleted_finders = make([]int, 0)
				}
				deleted_finders = append(deleted_finders, f_ix)
			}
		}

		if deleted_finders != nil {
			slices.Reverse(deleted_finders)
			for _, f_ix := range deleted_finders {
				*finders = slices.Delete(*finders, f_ix, f_ix+1)
			}
		}
	}
	return d - 1
}
