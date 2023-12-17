package day10

import (
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {

	area := getArea("day10/input.txt")
	w_area := len(area[0])
	h_area := len(area)

	area_sides := make([][]rune, h_area)
	for ix := range area_sides {
		area_sides[ix] = make([]rune, w_area)
	}
	trySetAreaSides := func(x, y int, r rune) {
		if x < 0 || y < 0 || x >= w_area || y >= h_area {
			return
		}
		area_sides[y][x] = r
	}

	fp_loop := '#'
	left := 'A'
	right := 'B'

	finders := getInitialFinders(area, 'S')
	turns := 0

	stepWithTurns := func(f_ix int, x int, y int, look bool) {

		if !look {
			return
		}
		f := finders[f_ix]
		next_r := area[y][x]
		area[y][x] = fp_loop

		curr_str := string(f.r)
		_ = curr_str
		next_str := string(next_r)
		_ = next_str

		if pointsUp(f.r, true) && f.y > y {

			if f.r == 'J' {
				trySetAreaSides(f.x, f.y+1, right)
				trySetAreaSides(f.x+1, f.y, right)
			} else if f.r == '|' {
				trySetAreaSides(f.x-1, f.y, left)
				trySetAreaSides(f.x+1, f.y, right)
			} else if f.r == 'L' {
				trySetAreaSides(f.x-1, f.y, left)
				trySetAreaSides(f.x, f.y+1, left)
			}

			if pointsLeft(next_r, false) {
				turns--
			} else if pointsRight(next_r, false) {
				turns++
			}

		} else if pointsDown(f.r, true) && f.y < y {

			if f.r == '7' {
				trySetAreaSides(f.x, f.y-1, left)
				trySetAreaSides(f.x+1, f.y, left)
			} else if f.r == '|' {
				trySetAreaSides(f.x-1, f.y, right)
				trySetAreaSides(f.x+1, f.y, left)
			} else if f.r == 'F' {
				trySetAreaSides(f.x, f.y-1, right)
				trySetAreaSides(f.x-1, f.y, right)
			}

			if pointsLeft(next_r, false) {
				turns++
			} else if pointsRight(next_r, false) {
				turns--
			}

		} else if pointsLeft(f.r, true) && f.x > x {

			if f.r == 'J' {
				trySetAreaSides(f.x, f.y+1, left)
				trySetAreaSides(f.x+1, f.y, left)
			} else if f.r == '-' {
				trySetAreaSides(f.x, f.y-1, right)
				trySetAreaSides(f.x, f.y+1, left)
			} else if f.r == '7' {
				trySetAreaSides(f.x, f.y-1, right)
				trySetAreaSides(f.x+1, f.y, right)
			}

			if pointsUp(next_r, false) {
				turns++
			} else if pointsDown(next_r, false) {
				turns--
			}

		} else if pointsRight(f.r, true) && f.x < x {

			if f.r == 'L' {
				trySetAreaSides(f.x-1, f.y, right)
				trySetAreaSides(f.x, f.y+1, right)
			} else if f.r == '-' {
				trySetAreaSides(f.x, f.y-1, left)
				trySetAreaSides(f.x, f.y+1, right)
			} else if f.r == 'F' {
				trySetAreaSides(f.x-1, f.y, left)
				trySetAreaSides(f.x, f.y-1, left)
			}

			if pointsUp(next_r, false) {
				turns--
			} else if pointsDown(next_r, false) {
				turns++
			}
		}

		f.x = x
		f.y = y
		f.r = next_r
		finders[f_ix] = f
	}

	runAssignment1(area, &finders, fp_loop, stepWithTurns)

	fp_inside := left
	if turns > 0 {
		fp_inside = right
	}

	finders = make([]finder, 0)
	for y, line := range area {
		for x, r_area := range line {
			if r_area == fp_loop {
				continue
			}
			if area_sides[y][x] == fp_inside {
				area[y][x] = fp_inside
				finders = append(finders, finder{x, y, fp_inside})
			}
		}
	}

	for _, line := range area_sides {
		for ix, r := range line {
			if r == '\x00' {
				line[ix] = ' '
			}
		}
	}

	canGoUp := func(f finder, _ bool) bool {
		r := area[f.y-1][f.x]
		return r != fp_inside && r != fp_loop
	}
	canGoDown := func(f finder, _ bool) bool {
		r := area[f.y+1][f.x]
		return r != fp_inside && r != fp_loop
	}
	canGoLeft := func(f finder, _ bool) bool {
		r := area[f.y][f.x-1]
		return r != fp_inside && r != fp_loop
	}
	canGoRight := func(f finder, _ bool) bool {
		r := area[f.y][f.x+1]
		return r != fp_inside && r != fp_loop
	}

	step := buildStep(area, &finders, fp_inside)
	runFinders(area, &finders, fp_inside, step,
		canGoUp,
		canGoDown,
		canGoLeft,
		canGoRight)

	inside := 0
	for _, line := range area {
		for _, r := range line {
			if r == fp_inside {
				inside++
			}
		}
	}

	clipboard.WriteAll(strconv.Itoa(inside))
}
