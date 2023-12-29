package main

import (
	"aoc2023/day1"
	"aoc2023/day10"
	"aoc2023/day11"
	"aoc2023/day12"
	"aoc2023/day13"
	"aoc2023/day14"
	"aoc2023/day15"
	"aoc2023/day16"
	"aoc2023/day17"
	"aoc2023/day18"
	"aoc2023/day19"
	"aoc2023/day2"
	"aoc2023/day20"
	"aoc2023/day3"
	"aoc2023/day4"
	"aoc2023/day5"
	"aoc2023/day6"
	"aoc2023/day7"
	"aoc2023/day8"
	"aoc2023/day9"
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
	fmt.Printf("Solving day %d assignment %d...\n", day, a)
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
	case 3:
		switch a {
		case 1:
			day3.Assignment1()
		case 2:
			day3.Assignment2()
		}
	case 4:
		switch a {
		case 1:
			day4.Assignment1()
		case 2:
			day4.Assignment2()
		}
	case 5:
		switch a {
		case 1:
			day5.Assignment1()
		case 2:
			day5.Assignment2()
		}
	case 6:
		switch a {
		case 1:
			day6.Assignment1()
		case 2:
			day6.Assignment2()
		}
	case 7:
		switch a {
		case 1:
			day7.Assignment1()
		case 2:
			day7.Assignment2()
		}
	case 8:
		switch a {
		case 1:
			day8.Assignment1()
		case 2:
			day8.Assignment2()
		}
	case 9:
		switch a {
		case 1:
			day9.Assignment1()
		case 2:
			day9.Assignment2()
		}
	case 10:
		switch a {
		case 1:
			day10.Assignment1()
		case 2:
			day10.Assignment2()
		}
	case 11:
		switch a {
		case 1:
			day11.Assignment1()
		case 2:
			day11.Assignment2()
		}
	case 12:
		switch a {
		case 1:
			day12.Assignment1()
		case 2:
			day12.Assignment2()
		}
	case 13:
		switch a {
		case 1:
			day13.Assignment1()
		case 2:
			day13.Assignment2()
		}
	case 14:
		switch a {
		case 1:
			day14.Assignment1()
		case 2:
			day14.Assignment2()
		}
	case 15:
		switch a {
		case 1:
			day15.Assignment1()
		case 2:
			day15.Assignment2()
		}
	case 16:
		switch a {
		case 1:
			day16.Assignment1()
		case 2:
			day16.Assignment2()
		}
	case 17:
		switch a {
		case 1:
			day17.Assignment1()
		case 2:
			day17.Assignment2()
		}
	case 18:
		switch a {
		case 1:
			day18.Assignment1()
		case 2:
			day18.Assignment2()
		}
	case 19:
		switch a {
		case 1:
			day19.Assignment1()
		case 2:
			day19.Assignment2()
		}
	case 20:
		switch a {
		case 1:
			day20.Assignment1()
		case 2:
			day20.Assignment2()
		}
	default:
		fmt.Fprintln(os.Stderr, "Day not yet solved.")
		os.Exit(1)
	}
	fmt.Print("Solved!")
}
