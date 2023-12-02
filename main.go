package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	day, err := strconv.Atoi(args[0])
	if err != nil || day < 1 || day > 24 {
		fmt.Fprintln(os.Stderr, "Invalid day.")
		os.Exit(1)
	}
	a, _ := strconv.Atoi(args[1])
	if err != nil || a != 1 && a != 2 {
		fmt.Fprintln(os.Stderr, "Invalid assignment.")
		os.Exit(1)
	}
	fmt.Printf("Solving day %d assignment %d...%s", day, a, utils.Constants.Newline)
	switch day {
	case 1:
		switch a {
		case 1:
			day1.Assignment1()
		case 2:
			day1.Assignment2()
		}
	case 2:
		switch a {
		case 1:
			day2.Assignment1()
		case 2:
			day2.Assignment2()
		}
	default:
		fmt.Fprintln(os.Stderr, "Day not yet solved.")
		os.Exit(1)
	}
	fmt.Print("Solved!")
}
