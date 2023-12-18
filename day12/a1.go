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

func arrangements(r0_ix int, s_ix int, record []rune, sizes []int, cache map[key]int) int {
	size := sizes[s_ix]
	result := 0

	so_far := 0
	restrict_start, restrict_end := false, false

	for r_ix := r0_ix; r_ix < len(record); r_ix++ {

		if restrict_start && so_far == 0 {
			break
		}

		r := record[r_ix]

		if r == '.' {
			if restrict_end {
				break
			}
			so_far = 0
			continue
		}

		if r == '#' {
			restrict_end = true
			if so_far == 0 {
				restrict_start = true
			}
		}

		so_far++

		if so_far == size {
			last_ix := r_ix
			r_ix -= so_far - 1
			so_far = 0

			if last_ix+1 < len(record) && record[last_ix+1] == '#' {
				if record[last_ix-size+1] == '#' {
					break
				}
				continue
			}

			if s_ix == len(sizes)-1 {
				if slices.Index(record[last_ix+1:], '#') == -1 {
					result++
				}
			} else if last_ix+2 < len(record) {
				next_r_ix := last_ix + 2
				next_s_ix := s_ix + 1
				k := key{next_r_ix, next_s_ix}
				if v, found := cache[k]; found {
					result += v
				} else {
					v := arrangements(next_r_ix, next_s_ix, record, sizes, cache)
					cache[k] = v
					result += v
				}
			}
		}
	}

	return result
}

type key struct {
	r_ix int
	s_ix int
}
