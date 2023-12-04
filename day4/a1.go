package day4

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

type matcher struct {
	winning map[string]struct{}
	pipeIx  int
	colonIx int
}

func newMatcher() *matcher {
	return &matcher{nil, -1, -1}
}

func (m *matcher) Match(line string) int {
	if m.colonIx == -1 {
		m.colonIx = strings.IndexRune(line, ':')
	}
	numbers := strings.Split(line[m.colonIx+1:], " ")
	numbers = slices.DeleteFunc(numbers, func(v string) bool {
		return v == ""
	})
	if m.pipeIx == -1 {
		m.pipeIx = slices.IndexFunc(numbers, func(v string) bool {
			return v == "|"
		})
		m.winning = make(map[string]struct{}, len(numbers)-m.pipeIx)
	}
	for _, w := range numbers[:m.pipeIx] {
		m.winning[w] = struct{}{}
	}
	count := 0
	for _, h := range numbers[m.pipeIx+1:] {
		if _, found := m.winning[h]; found {
			count++
		}
	}
	clear(m.winning)
	return count
}

func Assignment1() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	worth := 0
	m := newMatcher()
	for scanner.Scan() {
		line := scanner.Text()
		count := m.Match(line)
		worth += int(math.Pow(2, float64(count-1)))
	}
	clipboard.WriteAll(strconv.Itoa(worth))
}
