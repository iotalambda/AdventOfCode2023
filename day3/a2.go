package day3

import (
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	parts, numbers := getPartsAndNumbers("day3/input.txt")

	totalRatio := 0
	for ploc, part := range parts {
		if part != '*' {
			continue
		}

		ratioParts := make(map[*int]struct{}, 0)
		for _, s := range ploc.Surroundings() {
			if n, found := numbers[s]; found {
				ratioParts[n] = struct{}{}
			}
		}

		if len(ratioParts) != 2 {
			continue
		}

		ratio := 1
		for rp := range ratioParts {
			ratio *= *rp
		}

		totalRatio += ratio
	}

	clipboard.WriteAll(strconv.Itoa(totalRatio))
}
