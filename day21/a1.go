package day21

import (
	"aoc2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	g := load("day21/input.txt")
	yS, xS := g.find('S')
	(*g)[yS][xS] = '.'
	s := g.runAssignment1(yS, xS, 64)
	clipboard.WriteAll(strconv.Itoa(s))
}

func (g *garden) runAssignment1(yS, xS, dMax int) int {
	psNext := []painter{{yS, xS, 0}}
	tryPaint := func(p *painter, yOff, xOff int) {
		if p2, done := p.tryPaint(g, yOff, xOff, dMax, 'X'); done {
			psNext = append(psNext, *p2)
		}
	}
	s, p := 0, dMax%2 == 0
	for d := 0; len(psNext) > 0; d++ {
		ps := psNext
		psNext = []painter{}
		for pIx := range ps {
			p := &ps[pIx]
			tryPaint(p, -1, 0)
			tryPaint(p, 1, 0)
			tryPaint(p, 0, -1)
			tryPaint(p, 0, 1)
		}
		if (p && (d+1)%2 == 0) || (!p && (d+1)%2 == 1) {
			s += len(psNext)
		}

	}
	g.reset()
	return s
}

type garden [][]rune

func load(path string) *garden {
	var g garden
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	utils.ForEachMatrix(scanner, func(matrix [][]rune) {
		g = matrix
	})
	return &g
}

func (g *garden) find(r rune) (y, x int) {
	for y, row := range *g {
		for x, r2 := range row {
			if r == r2 {
				return y, x
			}
		}
	}
	return -1, -1
}

func (g *garden) tryGet(y, x int) (r rune, found bool) {
	if y < 0 || y >= len(*g) || x < 0 || x >= len(*g) {
		return '?', false
	}

	return (*g)[y][x], true
}

type painter struct {
	y, x, d int
}

func (p *painter) tryPaint(g *garden, yOff, xOff, maxD int, c rune) (*painter, bool) {
	if p.d < maxD {
		if t, found := g.tryGet(p.y+yOff, p.x+xOff); found && t == '.' {
			p2 := *p
			p2.y += yOff
			p2.x += xOff
			p2.d++
			(*g)[p2.y][p2.x] = c
			return &p2, true
		}
	}
	return nil, false
}

func (g *garden) reset() {
	for y, row := range *g {
		for x, r := range row {
			if r == 'X' {
				(*g)[y][x] = '.'
			}
		}
	}
}
