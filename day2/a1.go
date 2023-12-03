package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment1() {

	allowed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	possibleGames := 0
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

		possible := true
		for color, ra := range required {
			aa, found := allowed[color]
			if !found || ra > aa {
				possible = false
				break
			}
		}
		if possible {
			game, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not parse game ID:", err)
				os.Exit(1)
			}
			possibleGames += game
		}
	}

	clipboard.WriteAll(strconv.Itoa(possibleGames))
}
