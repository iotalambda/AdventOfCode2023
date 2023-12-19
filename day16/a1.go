package day16

import (
	"aoc2023/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	contraption := contraption("day16/input.txt")
	total := runAssignment1(contraption, beam{0, -1, Right})
	clipboard.WriteAll(strconv.Itoa(total))
}

func contraption(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var contraption [][]rune
	utils.ForEachMatrix(scanner, func(matrix [][]rune) {
		contraption = matrix
	})
	return contraption
}

func runAssignment1(contraption [][]rune, initial_beam beam) int {
	h, w := len(contraption), len(contraption[0])
	energization := make([][]uint8, h)
	for y := range energization {
		energization[y] = make([]uint8, w)
	}

	contraptionAt := func(y, x int) (r rune, found bool) {
		if y >= h || y < 0 || x >= w || x < 0 {
			return '?', false
		}

		return contraption[y][x], true
	}

	beams := []beam{initial_beam}
	to_remove := make([]int, 0)
	to_add := make([]beam, 0)

	proceed := func(b0 beam, b_ix int, fw_y_offset int, fw_x_offset int, fw dir, splitter rune, slash_dir dir, bslash_dir dir) {
		r, found := contraptionAt(b0.y+fw_y_offset, b0.x+fw_x_offset)
		if !found {
			to_remove = append(to_remove, b_ix)
			return
		}

		e := &energization[b0.y+fw_y_offset][b0.x+fw_x_offset]
		if *e&uint8(fw) > 0 {
			to_remove = append(to_remove, b_ix)
			return
		}
		*e |= uint8(fw)

		b := &beams[b_ix]
		b.y += fw_y_offset
		b.x += fw_x_offset

		switch r {
		case '/':
			b.d = slash_dir
		case '\\':
			b.d = bslash_dir
		case splitter:
			*e |= uint8(All)
			b.d = slash_dir
			to_add = append(to_add, beam{b.y, b.x, bslash_dir})
		}
	}

	for {
		if len(to_remove) > 0 {
			slices.Reverse(to_remove)
			for _, b_ix := range to_remove {
				beams = slices.Delete(beams, b_ix, b_ix+1)
			}
			to_remove = make([]int, 0)
		}

		if len(to_add) > 0 {
			beams = append(beams, to_add...)
			to_add = make([]beam, 0)
		}

		if len(beams) == 0 {
			break
		}

		for b_ix, b := range beams {
			if b.d&Up > 0 {
				proceed(b, b_ix, -1, 0, Up, '-', Right, Left)
			} else if b.d&Down > 0 {
				proceed(b, b_ix, 1, 0, Down, '-', Left, Right)
			} else if b.d&Left > 0 {
				proceed(b, b_ix, 0, -1, Left, '|', Down, Up)
			} else if b.d&Right > 0 {
				proceed(b, b_ix, 0, 1, Right, '|', Up, Down)
			}
		}
	}

	total := 0
	for _, row := range energization {
		for _, v := range row {
			if v > 0 {
				total++
			}
		}
	}

	return total
}

type beam struct {
	y int
	x int
	d dir
}

type dir uint8

const (
	Up    dir = 1
	Left  dir = 2
	Right dir = 4
	Down  dir = 8
	All   dir = 1 + 2 + 4 + 8
)
