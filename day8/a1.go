package day8

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
	file, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var instructions []rune
	letters := int('Z' - 'A' + 1)
	network := make([][2]int, int(math.Pow(float64(letters), 3)))

	for l := 0; scanner.Scan(); l++ {
		if l == 0 {
			instructions = []rune(scanner.Text())
			continue
		}

		if l == 1 {
			continue
		}

		line := strings.Fields(scanner.Text())
		node := int(line[0][0]-'A') + int(line[0][1]-'A')*letters + int(line[0][2]-'A')*letters*letters
		l := int(line[2][1]-'A') + int(line[2][2]-'A')*letters + int(line[2][3]-'A')*letters*letters
		r := int(line[3][0]-'A') + int(line[3][1]-'A')*letters + int(line[3][2]-'A')*letters*letters
		network[node] = [...]int{l, r}
	}

	mod := len(instructions)
	goal := len(network) - 1
	node := network[0]
	for i := 0; ; i++ {
		var next int
		if instructions[i%mod] == 'L' {
			next = node[0]
		} else {
			next = node[1]
		}
		if next == goal {
			clipboard.WriteAll(strconv.Itoa(i + 1))
			break
		}
		node = network[next]
	}
}
