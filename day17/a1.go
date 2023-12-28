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

func runAssignment1(path string, minStreak, maxStreak int) {
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
				b := block{int(r - '0'), make([]int, maxStreak), make([]int, maxStreak), make([]int, maxStreak), make([]int, maxStreak)}
				for s := 0; s < maxStreak; s++ {
					b.minHlUp[s] = math.MaxInt
					b.minHlDown[s] = math.MaxInt
					b.minHlLeft[s] = math.MaxInt
					b.minHlRight[s] = math.MaxInt
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
	nextTips := []tip{}
	tryAddTip := func(t tip) {
		b, found := tryGetBlock(t.y, t.x)
		if !found {
			return
		}
		add := false
		for d := t.step - 1; d < t.fuel; d++ {
			if mhl := b.getMinHeatLoss(t.dir)[d]; mhl > t.hlTotal {
				b.setMinHeatLoss(t.dir, d, t.hlTotal)
				add = true
			}
		}
		if !add {
			return
		}
		nextTips = append(nextTips, t)
	}
	tryAddTip(tip{0, 0, Right, minStreak, maxStreak, 0})
	tryAddTip(tip{0, 0, Down, minStreak, maxStreak, 0})

	for len(nextTips) > 0 {

		tips = nextTips
		nextTips = []tip{}

	tipsLoop:
		for _, t := range tips {
			offsetY, offsetX := offsets(t.dir)
			lh, rh := sides(t.dir)
			hlStreak := t.hlTotal
			y, x := t.y, t.x
			for s := 0; s < t.step; s++ {
				y += offsetY
				x += offsetX
				b, found := tryGetBlock(y, x)
				if !found {
					continue tipsLoop
				}
				hlStreak += b.hl
				t.fuel--
			}
			tryAddTip(tip{y, x, lh, minStreak, maxStreak, hlStreak})
			tryAddTip(tip{y, x, rh, minStreak, maxStreak, hlStreak})
			tryAddTip(tip{y, x, t.dir, 1, t.fuel, hlStreak})
		}
	}

	goal := city[h-1][w-1]
	result := slices.Min(goal.minHlDown)
	result = min(result, slices.Min(goal.minHlRight))
	clipboard.WriteAll(strconv.Itoa(result))
}

type block struct {
	hl         int
	minHlUp    []int
	minHlDown  []int
	minHlLeft  []int
	minHlRight []int
}

func (b *block) getMinHeatLoss(dir rune) []int {
	switch dir {
	case Up:
		return b.minHlUp
	case Down:
		return b.minHlDown
	case Left:
		return b.minHlLeft
	case Right:
		return b.minHlRight
	}
	panic(nil)
}

func (b *block) setMinHeatLoss(dir rune, ix int, value int) {
	switch dir {
	case Up:
		b.minHlUp[ix] = value
	case Down:
		b.minHlDown[ix] = value
	case Left:
		b.minHlLeft[ix] = value
	case Right:
		b.minHlRight[ix] = value
	}
}

type tip struct {
	y, x    int
	dir     rune
	step    int
	fuel    int
	hlTotal int
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
