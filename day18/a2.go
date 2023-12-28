package day18

import (
	"fmt"
	"os"
	"strconv"
)

func Assignment2() {
	runAssignment1("day18/input.txt", func(instructions []instruction) {
		for ix := range instructions {
			i := &instructions[ix]
			dist, err := strconv.ParseInt(i.Hex[2:7], 16, 0)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not parse hex:", err)
				os.Exit(1)
			}
			var dir rune
			switch i.Hex[7] {
			case '0':
				dir = 'R'
			case '1':
				dir = 'D'
			case '2':
				dir = 'L'
			case '3':
				dir = 'U'
			default:
				panic(nil)
			}
			i.Dir = dir
			i.Dist = int(dist)
		}
	})
}
