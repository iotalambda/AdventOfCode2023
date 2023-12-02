package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	digitToEnglishStr := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
	digitToEnglish := map[int][]rune{}
	digitToHsilgne := map[int][]rune{}
	for k, v := range digitToEnglishStr {
		digitToEnglish[k] = []rune(v)
		digitToHsilgne[k] = slices.Clone(digitToEnglish[k])
		slices.Reverse(digitToHsilgne[k])
	}

	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		senur := slices.Clone(runes)
		slices.Reverse(senur)
		sum += 10*findDigit(runes, digitToEnglish) + findDigit(senur, digitToHsilgne)
	}
	clipboard.WriteAll(strconv.Itoa(sum))
}

func findDigit(line []rune, translations map[int][]rune) int {
	for i := 0; i < len(line); i++ {
		c := line[i]
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
		for k, v := range translations {
			if slices.Equal(v, line[i:min(i+len(v), len(line))]) {
				return k
			}
		}
	}
	fmt.Fprintln(os.Stderr, "No digits found on line:", string(line))
	os.Exit(1)
	panic(nil)
}
