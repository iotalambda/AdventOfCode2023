package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment2() {
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
		slices.Reverse(readings)
		sum += extrapolate(readings)
	}

	clipboard.WriteAll(strconv.Itoa(sum))
}
