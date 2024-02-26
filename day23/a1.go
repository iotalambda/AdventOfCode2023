package day23

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	runAssignment1("day23/input.txt", true)
}

func runAssignment1(path string, slippery bool) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	t := make(trails, 0)
	nodes := make(map[location]*node, 0)
	for scanner.Scan() {
		t = append(t, []rune(scanner.Text()))
	}
	for y, r := range t {
		for x, r := range r {
			if r != '#' {
				c := 0
				if r, f := t.tryGet(y-1, x); f && r != '#' {
					c++
				}
				if r, f := t.tryGet(y, x-1); f && r != '#' {
					c++
				}
				if r, f := t.tryGet(y, x+1); f && r != '#' {
					c++
				}
				if r, f := t.tryGet(y+1, x); f && r != '#' {
					c++
				}
				if c > 2 {
					nodes[location{y, x}] = &node{map[*node]edge{}, false}
				}
			}
		}
	}

	var start, end *node
	var lStart, lEnd location
	for x, r := range t[0] {
		if r == '.' {
			lStart = location{0, x}
			start = &node{map[*node]edge{}, false}
			nodes[lStart] = start
			break
		}
	}
	for x, r := range t[len(t)-1] {
		if r == '.' {
			lEnd = location{len(t) - 1, x}
			end = &node{map[*node]edge{}, false}
			nodes[lEnd] = end
			break
		}
	}

	connect := func(origin *node, lOrigin location, y, x int) {
		history := make(map[location]struct{})
		history[lOrigin] = struct{}{}
		test := func(l location, b rune) bool {
			if _, f := history[l]; f {
				return false
			}
			if r, f := t.tryGet(l.y, l.x); f && r != '#' && (!slippery || r != b) {
				y, x = l.y, l.x
				return true
			}
			return false
		}

		path := make([]location, 0)
		for d := 1; ; d++ {
			l := location{y, x}
			path = append(path, l)
			if target, f := nodes[l]; f {
				if origin != target {
					origin.edges[target] = edge{path, d}
				}
				break
			}
			history[l] = struct{}{}
			if test(location{y - 1, x}, 'v') ||
				test(location{y, x - 1}, '>') ||
				test(location{y, x + 1}, '<') ||
				test(location{y + 1, x}, '^') {
				continue
			}
			break
		}
	}

	for l, n := range nodes {
		if r, f := t.tryGet(l.y-1, l.x); f && r != '#' && (!slippery || r != 'v') {
			connect(n, l, l.y-1, l.x)
		}
		if r, f := t.tryGet(l.y, l.x-1); f && r != '#' && (!slippery || r != '>') {
			connect(n, l, l.y, l.x-1)
		}
		if r, f := t.tryGet(l.y, l.x+1); f && r != '#' && (!slippery || r != '<') {
			connect(n, l, l.y, l.x+1)
		}
		if r, f := t.tryGet(l.y+1, l.x); f && r != '#' && (!slippery || r != '^') {
			connect(n, l, l.y+1, l.x)
		}
	}

	var getDMax func(n *node) int
	getDMax = func(n *node) int {
		if n == end {
			return 0
		}

		m := -1
		for n, e := range n.edges {
			if n.visited {
				continue
			}
			n.visited = true
			if dMax := getDMax(n); dMax != -1 {
				m = max(m, e.distance+dMax)
			}
			n.visited = false
		}

		return m
	}

	dMax := getDMax(start)
	clipboard.WriteAll(strconv.Itoa(dMax))
}

type trails [][]rune

type location struct {
	y, x int
}

type node struct {
	edges   map[*node]edge
	visited bool
}

type edge struct {
	path     []location
	distance int
}

func (t trails) tryGet(y, x int) (rune, bool) {
	if y < 0 || x < 0 || y >= len(t) || x >= len(t[0]) {
		return '\x00', false
	}
	return t[y][x], true
}
