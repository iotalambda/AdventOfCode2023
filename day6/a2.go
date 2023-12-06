package day6

import (
	"strings"
)

func Assignment2() {
	runDay6(func(lines *[2][]string) {
		for i, line := range lines {
			lines[i] = []string{strings.Join(line, "")}
		}
	})
}
