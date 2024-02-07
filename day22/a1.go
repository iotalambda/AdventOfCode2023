package day22

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
	runAssignment1("day22/input.txt", func(w world, bs []*brick) {
		safe := len(bs)
		for _, bR := range bs {
			w.removeBrick(bR)
			for _, bD := range bs {
				if d := w.getDropDistance(bD); d > 0 {
					safe--
					break
				}
			}
			w.addBrick(bR)
		}
		clipboard.WriteAll(strconv.Itoa(safe))
	})
}

func runAssignment1(path string, processDropped func(w world, bs []*brick)) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var zMax, yMax, xMax int
	bs := make([]*brick, 0)
	for scanner.Scan() {
		line := scanner.Text()
		ends := strings.Split(line, "~")
		s := utils.Ints(strings.Split(ends[0], ","))
		e := utils.Ints(strings.Split(ends[1], ","))
		b := brick{s[2], s[1], s[0], e[2], e[1], e[0]}
		zMax = max(zMax, b.z1, b.z2)
		yMax = max(yMax, b.y1, b.y2)
		xMax = max(xMax, b.x1, b.x2)
		bs = append(bs, &b)
	}
	slices.SortFunc(bs, func(a, b *brick) int {
		return min(a.z1, a.z2) - min(b.z1, b.z2)
	})

	w := newWorld(zMax, yMax, xMax)
	for _, b := range bs {
		w.dropBrick(b, w.getDropDistance(b))
		w.addBrick(b)
	}

	processDropped(w, bs)
}

type brick struct {
	z1, y1, x1, z2, y2, x2 int
}

func (b *brick) forEachZYX(action func(z, y, x int) bool) {
	zMin, zMax := min(b.z1, b.z2), max(b.z1, b.z2)
	yMin, yMax := min(b.y1, b.y2), max(b.y1, b.y2)
	xMin, xMax := min(b.x1, b.x2), max(b.x1, b.x2)
loop:
	for z := zMin; z <= zMax; z++ {
		for y := yMin; y <= yMax; y++ {
			for x := xMin; x <= xMax; x++ {
				if !action(z, y, x) {
					break loop
				}
			}
		}
	}
}

func (b *brick) forEachYX(action func(y, x int) bool) {
	yMin, yMax := min(b.y1, b.y2), max(b.y1, b.y2)
	xMin, xMax := min(b.x1, b.x2), max(b.x1, b.x2)
loop:
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if !action(y, x) {
				break loop
			}
		}
	}
}

type world [][][]*brick

func newWorld(zMax, yMax, xMax int) world {
	w := make(world, zMax+1)
	for z := range w {
		w[z] = make([][]*brick, yMax+1)
		for y := range w[z] {
			w[z][y] = make([]*brick, xMax+1)
		}
	}
	return w
}

func (w world) addBrick(b *brick) {
	b.forEachZYX(func(z, y, x int) bool {
		w[z][y][x] = b
		return true
	})
}

func (w world) removeBrick(b *brick) {
	b.forEachZYX(func(z, y, x int) bool {
		w[z][y][x] = nil
		return true
	})
}

func (w world) dropBrick(b *brick, d int) {
	b.z1 -= d
	b.z2 -= d
}

func (w world) getDropDistance(b *brick) int {
	d := 0
	collision := false
	for z := min(b.z1, b.z2); z > 0; z-- {
		b.forEachYX(func(yB, xB int) bool {
			if w[z-1][yB][xB] != nil {
				collision = true
				return false
			}
			return true
		})

		if collision {
			break
		}

		d++
	}

	return d
}
