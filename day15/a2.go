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
	label_box_ixs := make(map[string]int, 0)

	boxes := make([][]lens, 256)
	for ix := range boxes {
		boxes[ix] = make([]lens, 0)
	}

	for _, part := range sequence {

		op_ix := strings.IndexRune(part, '=')
		if op_ix == -1 {
			op_ix = strings.IndexRune(part, '-')
		}

		label := part[:op_ix]

		box_ix, found := label_box_ixs[label]
		if !found {
			box_ix = hash(label)
			label_box_ixs[label] = box_ix
		}
		box := boxes[box_ix]

		old_lens_ix := slices.IndexFunc(box, func(l lens) bool {
			return l.label == label
		})

		op := part[op_ix]
		switch op {
		case '=':
			focal, err := strconv.Atoi(part[op_ix+1:])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not parse focal:", err)
				os.Exit(1)
			}

			new_lens := lens{label, focal}

			if old_lens_ix == -1 {
				box = append(box, new_lens)
				boxes[box_ix] = box
			} else {
				box[old_lens_ix] = new_lens
			}
		case '-':
			if old_lens_ix != -1 {
				boxes[box_ix] = slices.Delete(box, old_lens_ix, old_lens_ix+1)
			}
		}
	}

	focusing_power := 0
	for box_ix, box := range boxes {
		for lens_ix, lens := range box {
			focusing_power += (box_ix + 1) * (lens_ix + 1) * lens.focal
		}
	}
	clipboard.WriteAll(strconv.Itoa(focusing_power))
}

type lens struct {
	label string
	focal int
}
