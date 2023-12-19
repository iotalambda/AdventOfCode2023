package day16

import (
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	max_total := 0
	beams := make([]beam, 0)
	contraption := contraption("day16/input.txt")
	h, w := len(contraption), len(contraption[0])
	for y := 0; y < h; y++ {
		beams = append(beams, beam{y, -1, Right}, beam{y, w, Left})
	}
	for x := 1; x < w-1; x++ {
		beams = append(beams, beam{-1, x, Down}, beam{h, x, Up})
	}
	for _, b := range beams {
		max_total = max(runAssignment1(contraption, b), max_total)
	}
	clipboard.WriteAll(strconv.Itoa(max_total))
}
