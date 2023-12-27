package day17

import (
	"aoc2023/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	runAssignment1("day17/input.txt", 1, 3)
}

func runAssignment1(path string, min_streak, max_streak int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var city [][]block
	utils.ForEachMatrix(scanner, func(matrix [][]rune) {
		city = make([][]block, len(matrix))
		for y, row := range matrix {
			city[y] = make([]block, len(row))
			for x, r := range row {
				b := block{int(r - '0'), make([]int, max_streak), make([]int, max_streak), make([]int, max_streak), make([]int, max_streak)}
				for s := 0; s < max_streak; s++ {
					b.min_hl_up[s] = math.MaxInt
					b.min_hl_down[s] = math.MaxInt
					b.min_hl_left[s] = math.MaxInt
					b.min_hl_right[s] = math.MaxInt
				}
				city[y][x] = b
			}
		}
	})
	h, w := len(city), len(city[0])
	tryGetBlock := func(y, x int) (*block, bool) {
		if y < 0 || y >= h || x < 0 || x >= w {
			return nil, false
		}
		return &city[y][x], true
	}

	var tips []tip
	next_tips := []tip{}
	tryAddTip := func(t tip) {
		b, found := tryGetBlock(t.y, t.x)
		if !found {
			return
		}
		add := false
		for d := t.step - 1; d < t.fuel; d++ {
			if mhl := b.getMinHeatLoss(t.dir)[d]; mhl > t.hl_total {
				b.setMinHeatLoss(t.dir, d, t.hl_total)
				add = true
			}
		}
		if !add {
			return
		}
		next_tips = append(next_tips, t)
	}
	tryAddTip(tip{0, 0, Right, min_streak, max_streak, 0})
	tryAddTip(tip{0, 0, Down, min_streak, max_streak, 0})

	for len(next_tips) > 0 {

		tips = next_tips
		next_tips = []tip{}

	tips_loop:
		for _, t := range tips {
			offset_y, offset_x := offsets(t.dir)
			lh, rh := sides(t.dir)
			hl_streak := t.hl_total
			y, x := t.y, t.x
			for s := 0; s < t.step; s++ {
				y += offset_y
				x += offset_x
				b, found := tryGetBlock(y, x)
				if !found {
					continue tips_loop
				}
				hl_streak += b.hl
				t.fuel--
			}
			tryAddTip(tip{y, x, lh, min_streak, max_streak, hl_streak})
			tryAddTip(tip{y, x, rh, min_streak, max_streak, hl_streak})
			tryAddTip(tip{y, x, t.dir, 1, t.fuel, hl_streak})
		}
	}

	goal := city[h-1][w-1]
	result := slices.Min(goal.min_hl_down)
	result = min(result, slices.Min(goal.min_hl_right))
	clipboard.WriteAll(strconv.Itoa(result))
}

type block struct {
	hl           int
	min_hl_up    []int
	min_hl_down  []int
	min_hl_left  []int
	min_hl_right []int
}

func (b *block) getMinHeatLoss(dir rune) []int {
	switch dir {
	case Up:
		return b.min_hl_up
	case Down:
		return b.min_hl_down
	case Left:
		return b.min_hl_left
	case Right:
		return b.min_hl_right
	}
	panic(nil)
}

func (b *block) setMinHeatLoss(dir rune, ix int, value int) {
	switch dir {
	case Up:
		b.min_hl_up[ix] = value
	case Down:
		b.min_hl_down[ix] = value
	case Left:
		b.min_hl_left[ix] = value
	case Right:
		b.min_hl_right[ix] = value
	}
}

type tip struct {
	y, x     int
	dir      rune
	step     int
	fuel     int
	hl_total int
}

const (
	Up    = 'U'
	Down  = 'D'
	Left  = 'L'
	Right = 'R'
)

func offsets(dir rune) (y int, x int) {
	switch dir {
	case Up:
		return -1, 0
	case Down:
		return 1, 0
	case Left:
		return 0, -1
	case Right:
		return 0, 1
	}
	panic(nil)
}

func sides(dir rune) (lh rune, rh rune) {
	switch dir {
	case Up:
		return Left, Right
	case Down:
		return Right, Left
	case Left:
		return Down, Up
	case Right:
		return Up, Down
	}
	panic(nil)
}
