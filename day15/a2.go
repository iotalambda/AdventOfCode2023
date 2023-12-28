package day15

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
	file, err := os.Open("day15/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	sequence := strings.Split(scanner.Text(), ",")
	labelBoxIxs := make(map[string]int, 0)

	boxes := make([][]lens, 256)
	for ix := range boxes {
		boxes[ix] = make([]lens, 0)
	}

	for _, part := range sequence {

		opIx := strings.IndexRune(part, '=')
		if opIx == -1 {
			opIx = strings.IndexRune(part, '-')
		}

		label := part[:opIx]

		boxIx, found := labelBoxIxs[label]
		if !found {
			boxIx = hash(label)
			labelBoxIxs[label] = boxIx
		}
		box := boxes[boxIx]

		oldLensIx := slices.IndexFunc(box, func(l lens) bool {
			return l.label == label
		})

		op := part[opIx]
		switch op {
		case '=':
			focal, err := strconv.Atoi(part[opIx+1:])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not parse focal:", err)
				os.Exit(1)
			}

			newLens := lens{label, focal}

			if oldLensIx == -1 {
				box = append(box, newLens)
				boxes[boxIx] = box
			} else {
				box[oldLensIx] = newLens
			}
		case '-':
			if oldLensIx != -1 {
				boxes[boxIx] = slices.Delete(box, oldLensIx, oldLensIx+1)
			}
		}
	}

	focusingPower := 0
	for boxIx, box := range boxes {
		for lensIx, lens := range box {
			focusingPower += (boxIx + 1) * (lensIx + 1) * lens.focal
		}
	}
	clipboard.WriteAll(strconv.Itoa(focusingPower))
}

type lens struct {
	label string
	focal int
}
