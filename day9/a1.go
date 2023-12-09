package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		readings := ints(strings.Fields(scanner.Text()))
		sum += extrapolate(readings)
	}

	clipboard.WriteAll(strconv.Itoa(sum))
}

func extrapolate(readings []int) int {
	go_deeper := false
	diffs := make([]int, len(readings)-1)
	for ix, r := range readings[1:] {
		diff := r - readings[ix]
		diffs[ix] = diff
		if diff != 0 {
			go_deeper = true
		}
	}
	if go_deeper {
		return readings[len(readings)-1] + extrapolate(diffs)
	}
	return readings[0]
}

func ints(source []string) []int {
	result := make([]int, len(source))
	for i, str := range source {
		v, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse int:", err)
			os.Exit(1)
		}
		result[i] = v
	}
	return result
}
