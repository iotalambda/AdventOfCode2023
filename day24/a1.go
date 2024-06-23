package day24

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
	hailstones := read("day24/input.txt")
	interactions := 0
	pMin, pMax := 200000000000000., 400000000000000.
	for h1Ix, h1 := range hailstones[:len(hailstones)-1] {
		for _, h2 := range hailstones[h1Ix+1:] {
			d := h1.v.y*h2.v.x - h2.v.y*h1.v.x
			if d == 0 {
				continue
			}
			n := (h2.p0.y-h1.p0.y)*(h1.v.x*h2.v.x) + h1.v.y*h2.v.x*h1.p0.x - h2.v.y*h1.v.x*h2.p0.x
			x := n / d
			k := h1.v.y * (x - h1.p0.x) / h1.v.x
			y := k + h1.p0.y

			if y < pMin || y > pMax || x < pMin || x > pMax {
				continue
			}
			if h1.p0.y < y && h1.v.y <= 0 ||
				h1.p0.y > y && h1.v.y >= 0 ||
				h1.p0.x < x && h1.v.x <= 0 ||
				h1.p0.x > x && h1.v.x >= 0 ||
				h2.p0.y < y && h2.v.y <= 0 ||
				h2.p0.y > y && h2.v.y >= 0 ||
				h2.p0.x < x && h2.v.x <= 0 ||
				h2.p0.x > x && h2.v.x >= 0 {
				continue
			}
			interactions++
		}
	}
	clipboard.WriteAll(strconv.Itoa(interactions))
}

type trajectory struct {
	p0 v3d
	v  v3d
}

type v3d struct {
	z, y, x float64
}

func read(path string) []trajectory {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	hailstones := make([]trajectory, 0)
	for scanner.Scan() {
		strs := strings.Split(strings.ReplaceAll(scanner.Text(), ",", ""), " ")
		strs = slices.DeleteFunc(strs, func(s string) bool {
			return s == "@" || s == ""
		})
		floats := make([]float64, 0)
		for _, i := range utils.Ints(strs) {
			floats = append(floats, float64(i))
		}
		hailstones = append(hailstones, trajectory{v3d{floats[2], floats[1], floats[0]}, v3d{floats[5], floats[4], floats[3]}})
	}
	return hailstones
}
