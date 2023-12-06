package day6

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	runDay6(func(lines *[2][]string) {})
}

func runDay6(prepareLines func(lines *[2][]string)) {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var linesStr [2][]string
	var lines [2][]float64
	li := 0
	for scanner.Scan() {
		linesStr[li] = strings.Fields(scanner.Text())[1:]
		li++
	}
	prepareLines(&linesStr)
	for i, lineStr := range linesStr {
		lines[i] = floats(lineStr)
	}

	margin := 1
	for i, t_total := range lines[0] {
		d_record := lines[1][i]
		part := math.Sqrt(math.Pow(t_total, 2) - 4*d_record)
		t_press_1 := (-t_total + part) / -2
		t_press_2 := (-t_total - part) / -2
		t_press_l := int(math.Floor(max(min(t_press_1, t_press_2), -1)) + 1)
		t_press_u := int(math.Ceil(max(t_press_1, t_press_2, 0)) - 1)
		margin *= t_press_u - t_press_l + 1
	}
	clipboard.WriteAll(strconv.Itoa(int(margin)))
}

func floats(strs []string) []float64 {
	result := make([]float64, len(strs))
	for i, str := range strs {
		r, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse int:", err)
			os.Exit(1)
		}
		result[i] = float64(r)
	}
	return result
}
