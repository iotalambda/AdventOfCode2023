package day19

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment1() {

	wfs, parts := loadAssignment("day19/input.txt")

	result := 0

	for _, p := range parts {
		wfN := "in"
	workflowLoop:
		for {
			switch wfN {
			case "A":
				result += p['x'] + p['m'] + p['a'] + p['s']
				fallthrough
			case "R":
				break workflowLoop
			default:
				wf := wfs[wfN]
				for _, r := range wf.Rs {
					if r.GT && p[r.C] > r.V || !r.GT && p[r.C] < r.V {
						wfN = r.WfN
						continue workflowLoop
					} else {
						continue
					}
				}
				wfN = wf.F
			}
		}
	}

	clipboard.WriteAll(strconv.Itoa(result))
}

func loadAssignment(path string) (map[string]workflow, []part) {
	wfs := map[string]workflow{}
	parts := []part{}

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	stage := "wfs"

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			stage = "parts"
			continue
		}
		switch stage {

		case "wfs":
			strs := strings.Split(strings.TrimRight(line, "}"), "{")
			n := strs[0]
			strs = strings.Split(strs[1], ",")
			f := strs[len(strs)-1]
			rs := []rule{}
			for _, strs := range strs[:len(strs)-1] {
				strs := strings.Split(strs, ":")
				wfN := strs[1]
				s := ">"
				if strings.Contains(strs[0], "<") {
					s = "<"
				}
				strs = strings.Split(strs[0], s)
				rs = append(rs, rule{category(strs[0][0]), s == ">", parseInt(strs[1]), wfN})
			}
			wfs[n] = workflow{rs, f}

		case "parts":
			cvsStr := strings.Split(strings.Trim(line, "{}"), ",")
			part := part{}
			for _, cv := range cvsStr {
				cv := strings.Split(cv, "=")
				part[category(cv[0][0])] = parseInt(cv[1])
			}
			parts = append(parts, part)
		}
	}

	return wfs, parts
}

func parseInt(src string) int {
	v, err := strconv.Atoi(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not parse int:", err)
		os.Exit(1)
	}
	return v
}

type workflow struct {
	Rs []rule
	F  string
}

type rule struct {
	C   category
	GT  bool
	V   int
	WfN string
}

type part map[category]int

type category rune
