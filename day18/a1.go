package day18

import (
	"aoc2023/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	runAssignment1("day18/input.txt", func(i []instruction) {})
}

func runAssignment1(path string, prepare func([]instruction)) {
	instructions := []instruction{}
	turns := 0

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	prevDir := '\x00'
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse int:", err)
			os.Exit(1)
		}
		dir := rune(parts[0][0])
		instructions = append(instructions, instruction{dir, distance, parts[2]})
		if prevDir == 'U' && dir == 'R' ||
			prevDir == 'R' && dir == 'D' ||
			prevDir == 'D' && dir == 'L' ||
			prevDir == 'L' && dir == 'U' {
			turns++
		} else {
			turns--
		}
		prevDir = rune(dir)
	}

	prepare(instructions)

	if dir := instructions[0].Dir; dir == 'U' || dir == 'D' {
		instructions = append(instructions, instructions[0])
	} else {
		instructions = utils.Prepend(instructions, instructions[len(instructions)-1])
	}

	lines := make([]line, 0)

	y, x := 0, 0
	var prevLine *line
	var prevD rune
	for _, i := range instructions {
		var yOffset, xOffset int
		switch i.Dir {
		case 'U':
			yOffset, xOffset = -1, 0
		case 'D':
			yOffset, xOffset = 1, 0
		case 'L':
			yOffset, xOffset = 0, -1
		case 'R':
			yOffset, xOffset = 0, 1
		}
		yNext, xNext := y+yOffset*i.Dist, x+xOffset*i.Dist

		switch i.Dir {
		case 'U':
			fallthrough
		case 'D':
			if prevLine != nil {
				if prevLine.Out == i.Dir {
					if prevLine.Dir == 'L' {
						prevLine.X1++
					} else {
						prevLine.X2--
					}
				}
			}
			prevD = i.Dir

		case 'L':
			var out rune
			if turns > 0 {
				out = 'D'
			} else {
				out = 'U'
			}
			if prevD != out {
				x--
			}
			lines = append(lines, line{y, xNext, x, out, i.Dir})
			prevLine = &lines[len(lines)-1]

		case 'R':
			var out rune
			if turns > 0 {
				out = 'U'
			} else {
				out = 'D'
			}
			if prevD != out {
				x++
			}
			lines = append(lines, line{y, x, xNext, out, i.Dir})
			prevLine = &lines[len(lines)-1]
		}

		y, x = yNext, xNext
	}

	slices.SortFunc(lines, func(a, b line) int {
		return a.Y - b.Y
	})

	area := 0

	for topLIx, topL := range lines {
		if topL.Out != 'U' {
			continue
		}

		for xStart := topL.X1; xStart <= topL.X2; {
			bottomLIx := topLIx + 1 + slices.IndexFunc(lines[topLIx+1:], func(l line) bool {
				return l.Out == 'D' && l.X1 <= xStart && xStart <= l.X2
			})
			bottomL := lines[bottomLIx]
			xEnd := min(topL.X2, bottomL.X2)

			minShadowingLIx := topLIx + 1
			for {
				shadowingLIx := slices.IndexFunc(lines[minShadowingLIx:bottomLIx], func(l line) bool {
					return l.Out == 'D' && xStart < l.X1 && l.X1 < xEnd
				})
				if shadowingLIx == -1 {
					break
				}
				shadowingLIx += minShadowingLIx
				shadowingL := lines[shadowingLIx]
				xEnd = shadowingL.X1 - 1
				minShadowingLIx = shadowingLIx + 1
			}

			area += (bottomL.Y - topL.Y + 1) * (xEnd - xStart + 1)
			xStart = xEnd + 1
		}
	}
	clipboard.WriteAll(strconv.Itoa(area))
}

type instruction struct {
	Dir  rune
	Dist int
	Hex  string
}

type line struct {
	Y, X1, X2 int
	Out, Dir  rune
}
