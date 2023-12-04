package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cards := make([]int, 0)
	addWins := func(cardIx int, amount int, multiplier int) {
		gap := max(0, cardIx+amount-len(cards)+1)
		for i := 0; i < gap; i++ {
			cards = append(cards, multiplier)
		}
		for ix := cardIx + 1; ix <= cardIx+amount-gap; ix++ {
			cards[ix] += multiplier
		}
	}
	m := newMatcher()
	cardIx := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := m.Match(line)
		var multiplier int
		if len(cards) > cardIx {
			cards[cardIx] += 1
			multiplier = cards[cardIx]
		} else {
			cards = append(cards, 1)
			multiplier = 1
		}
		addWins(cardIx, matches, multiplier)
		cardIx++
	}
	sum := 0
	for _, a := range cards {
		sum += a
	}
	clipboard.WriteAll(strconv.Itoa(sum))
}
