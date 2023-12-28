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

func runAssignment1(contraption [][]rune, initialBeam beam) int {
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

	beams := []beam{initialBeam}
	toRemove := make([]int, 0)
	toAdd := make([]beam, 0)

	proceed := func(b0 beam, bIx int, fwYOffset int, fwXOffset int, fw uint8, splitter rune, slashDir uint8, bslashDir uint8) {
		r, found := contraptionAt(b0.y+fwYOffset, b0.x+fwXOffset)
		if !found {
			toRemove = append(toRemove, bIx)
			return
		}

		e := &energization[b0.y+fwYOffset][b0.x+fwXOffset]
		if *e&fw > 0 {
			toRemove = append(toRemove, bIx)
			return
		}
		*e |= fw

		b := &beams[bIx]
		b.y += fwYOffset
		b.x += fwXOffset

		switch r {
		case '/':
			b.d = slashDir
		case '\\':
			b.d = bslashDir
		case splitter:
			*e |= All
			b.d = slashDir
			toAdd = append(toAdd, beam{b.y, b.x, bslashDir})
		}
	}

	for {
		if len(toRemove) > 0 {
			slices.Reverse(toRemove)
			for _, bIx := range toRemove {
				beams = slices.Delete(beams, bIx, bIx+1)
			}
			toRemove = make([]int, 0)
		}

		if len(toAdd) > 0 {
			beams = append(beams, toAdd...)
			toAdd = make([]beam, 0)
		}

		if len(beams) == 0 {
			break
		}

		for bIx, b := range beams {
			if b.d == Up {
				proceed(b, bIx, -1, 0, Up, '-', Right, Left)
			} else if b.d == Down {
				proceed(b, bIx, 1, 0, Down, '-', Left, Right)
			} else if b.d == Left {
				proceed(b, bIx, 0, -1, Left, '|', Down, Up)
			} else if b.d == Right {
				proceed(b, bIx, 0, 1, Right, '|', Up, Down)
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
	d uint8
}

const (
	Up    uint8 = 1
	Left  uint8 = 2
	Right uint8 = 4
	Down  uint8 = 8
	All   uint8 = Up + Down + Left + Right
)
