package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment2() {

	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalPower := 0
	for scanner.Scan() {

		parts := strings.Split(scanner.Text(), " ")
		for i, part := range parts {
			parts[i] = strings.TrimRightFunc(part, func(r rune) bool {
				return r == ':' || r == ',' || r == ';'
			})
		}
		required := make(map[string]int)
		start := 2
		for iAbsolute, part := range parts[start : len(parts)-1] {
			i := start + iAbsolute
			curr, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			color := parts[i+1]
			prev, found := required[color]
			if !found || curr > prev {
				required[color] = curr
			}
		}

		power := 1
		for _, a := range required {
			power *= a
		}
		totalPower += power
	}

	clipboard.WriteAll(strconv.Itoa(totalPower))
}
