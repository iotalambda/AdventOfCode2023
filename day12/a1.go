package day12

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
	runAssignment1("day12/input.txt", func(record []rune, sizes []int) ([]rune, []int) {
		return record, sizes
	})
}

func runAssignment1(path string, prepare func(record []rune, sizes []int) ([]rune, []int)) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		record := []rune(parts[0])
		sizes := utils.Ints(strings.Split(parts[1], ","))
		record, sizes = prepare(record, sizes)
		result := arrangements(0, 0, record, sizes, make(map[key]int))
		total += result
	}

	clipboard.WriteAll(strconv.Itoa(total))
}

func arrangements(r0Ix int, sIx int, record []rune, sizes []int, cache map[key]int) int {
	size := sizes[sIx]
	result := 0

	soFar := 0
	restrictStart, restrictEnd := false, false

	for rIx := r0Ix; rIx < len(record); rIx++ {

		if restrictStart && soFar == 0 {
			break
		}

		r := record[rIx]

		if r == '.' {
			if restrictEnd {
				break
			}
			soFar = 0
			continue
		}

		if r == '#' {
			restrictEnd = true
			if soFar == 0 {
				restrictStart = true
			}
		}

		soFar++

		if soFar == size {
			lastIx := rIx
			rIx -= soFar - 1
			soFar = 0

			if lastIx+1 < len(record) && record[lastIx+1] == '#' {
				if record[lastIx-size+1] == '#' {
					break
				}
				continue
			}

			if sIx == len(sizes)-1 {
				if slices.Index(record[lastIx+1:], '#') == -1 {
					result++
				}
			} else if lastIx+2 < len(record) {
				nextRIx := lastIx + 2
				nextSIx := sIx + 1
				k := key{nextRIx, nextSIx}
				if v, found := cache[k]; found {
					result += v
				} else {
					v := arrangements(nextRIx, nextSIx, record, sizes, cache)
					cache[k] = v
					result += v
				}
			}
		}
	}

	return result
}

type key struct {
	rIx int
	sIx int
}
