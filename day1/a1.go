package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var d1 int
		var d2 int
		for i := 0; i < len(line); i++ {
			c := rune(line[i])
			if unicode.IsDigit(c) {
				d1 = int(c - '0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			c := rune(line[i])
			if unicode.IsDigit(c) {
				d2 = int(c - '0')
				break
			}
		}
		sum += d1*10 + d2
	}
	clipboard.WriteAll(strconv.Itoa(sum))
}
