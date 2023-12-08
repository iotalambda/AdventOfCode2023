package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	n := 5
	cards := map[rune]int{
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'J': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}
	base := int(math.Pow(float64(len(cards)), float64(n)))
	hands := make([][2]int, 0)
	camounts := make([]int, len(cards))

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		bid, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse int:", err)
			os.Exit(1)
		}
		score := 0
		for i, r := range []rune(line[0]) {
			c := cards[r]
			score += c * int(math.Pow(float64(len(cards)), float64(n-i-1)))
			camounts[c] += 1
		}
		slices.SortFunc(camounts, func(a, b int) int {
			return b - a
		})
		switch camounts[0] {
		case 2:
			switch camounts[1] {
			case 2:
				score += base * 2
			case 1:
				score += base
			}
		case 3:
			switch camounts[1] {
			case 2:
				score += base * 4
			case 1:
				score += base * 3
			}
		case 4:
			score += base * 5
		case 5:
			score += base * 6
		}
		hands = append(hands, [...]int{score, bid})
		for i := range camounts {
			if camounts[i] == 0 {
				break
			}
			camounts[i] = 0
		}
	}

	slices.SortFunc(hands, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand[1]
	}

	clipboard.WriteAll(strconv.Itoa(total))
}
