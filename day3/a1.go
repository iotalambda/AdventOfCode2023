package day3

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"unicode"

	"github.com/atotto/clipboard"
)

type loc struct {
	X, Y int
}

func (l1 loc) Equal(l2 loc) bool {
	return l1.X == l2.X && l1.Y == l2.Y
}

func (l loc) Surroundings() [8]loc {
	return [8]loc{
		{l.X - 1, l.Y - 1},
		{l.X, l.Y - 1},
		{l.X + 1, l.Y - 1},
		{l.X - 1, l.Y},
		{l.X + 1, l.Y},
		{l.X - 1, l.Y + 1},
		{l.X, l.Y + 1},
		{l.X + 1, l.Y + 1},
	}
}

func getPartsAndNumbers(path string) (map[loc]rune, map[loc]*int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := make(map[loc]*int)
	parts := make(map[loc]rune)

	var digits []uint8
	saveNumber := func(x int, y int) {
		number := 0
		slices.Reverse(digits)
		for i, d := range digits {
			number += int(d) * int(math.Pow10(i))
			numbers[loc{x - i, y}] = &number
		}
		digits = nil
	}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			if unicode.IsDigit(c) {
				if digits == nil {
					digits = make([]uint8, 0)
				}
				digits = append(digits, uint8(c-'0'))
				if x == len(line)-1 {
					saveNumber(x, y)
				}
				continue
			}

			if digits != nil {
				saveNumber(x-1, y)
			}

			if c == '.' {
				continue
			}

			parts[loc{x, y}] = c
		}
		y++
	}

	return parts, numbers
}

func Assignment1() {

	parts, numbers := getPartsAndNumbers("day3/input.txt")

	partNumbers := make(map[*int]struct{}, 0)
	for ploc := range parts {
		for _, s := range ploc.Surroundings() {
			if n, found := numbers[s]; found {
				partNumbers[n] = struct{}{}
			}
		}
	}

	sum := 0
	for pn := range partNumbers {
		sum += *pn
	}

	clipboard.WriteAll(strconv.Itoa(sum))
}
