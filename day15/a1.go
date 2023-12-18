package day15

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	file, err := os.Open("day15/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sequence := strings.Split(scanner.Text(), ",")
	sum := 0
	for _, part := range sequence {
		sum += hash(part)
	}

	clipboard.WriteAll(strconv.Itoa(sum))
}

func hash(source string) int {
	current := 0
	for _, r := range source {
		current += int(r)
		current *= 17
		current %= 256
	}
	return current
}
