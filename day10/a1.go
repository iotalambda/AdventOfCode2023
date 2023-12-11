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
	file, err := os.Open("day10/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	area := make([][]rune, 0)

	up := uint8(1)
	down := uint8(2)
	left := uint8(4)
	right := uint8(8)

	isUp := func(r rune) bool {
		return r == 'J' || r == '|' || r == 'L'
	}
	isDown := func(r rune) bool {
		return r == '7' || r == '|' || r == 'F'
	}
	isLeft := func(r rune) bool {
		return r == 'L' || r == '-' || r == 'F'
	}
	isRight := func(r rune) bool {
		return r == 'J' || r == '-' || r == '7'
	}

	neighbours := func(f finder) uint8 {
		var result uint8
		if f.y > 0 && (isUp(f.r) || f.r == 'S') {
			if r := area[f.y-1][f.x]; isDown(r) {
				result |= up
			}
		}
		if f.y < len(area)-1 && (isDown(f.r) || f.r == 'S') {
			if r := area[f.y+1][f.x]; isUp(r) {
				result |= down
			}
		}
		if f.x > 0 && (isRight(f.r) || f.r == 'S') {
			if r := area[f.y][f.x-1]; isLeft(r) {
				result |= left
			}
		}
		if f.x < len(area[0])-1 && (isLeft(f.r) || f.r == 'S') {
			if r := area[f.y][f.x+1]; isRight(r) {
				result |= right
			}
		}
		return result
	}

	s_x, s_y := -1, -1
	for scanner.Scan() {
		line := []rune(scanner.Text())
		area = append(area, line)
		if s_x == -1 {
			s_x = slices.Index(line, 'S')
			s_y++
		}
	}

	finders := []finder{{s_x, s_y, 'S'}}

	step := func(f_ix int, x int, y int, n_ix *int) {
		next := finder{x, y, area[y][x]}
		if *n_ix == 0 {
			finders[f_ix] = next
		} else {
			finders = append(finders, next)
		}
		area[y][x] = '.'
		*n_ix++
	}

	var d int

	for ; len(finders) > 0; d++ {
		var deleted_finders []int
		for f_ix, f := range finders {
			ns := neighbours(f)

			if ns == 0 {
				if deleted_finders == nil {
					deleted_finders = make([]int, 0)
				}
				deleted_finders = append(deleted_finders, f_ix)
			}
			var n_ix int

			if (ns & up) == up {
				step(f_ix, f.x, f.y-1, &n_ix)
			}
			if (ns & down) == down {
				step(f_ix, f.x, f.y+1, &n_ix)
			}
			if (ns & left) == left {
				step(f_ix, f.x-1, f.y, &n_ix)
			}
			if (ns & right) == right {
				step(f_ix, f.x+1, f.y, &n_ix)
			}
		}

		if deleted_finders != nil {
			slices.Reverse(deleted_finders)
			for _, f_ix := range deleted_finders {
				finders = slices.Delete(finders, f_ix, f_ix+1)
			}
		}
	}
	d--

	clipboard.WriteAll(strconv.Itoa(d))
}

type finder struct {
	x int
	y int
	r rune
}
