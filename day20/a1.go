package day20

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	low, high, _ := runAssignment1("day20/input.txt", 1000, func(pulse) bool { return false })
	clipboard.WriteAll(strconv.Itoa(low * high))
}

func runAssignment1(path string, maxPresses int, stop func(pulse) bool) (low int, high int, presses int) {

	ms := map[string]module{}
	{
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not open file:", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		cs := []*conjuction{}
		for scanner.Scan() {
			parts := strings.Split(strings.Replace(scanner.Text(), " ", "", -1), "->")
			n := parts[0]
			parts = strings.Split(parts[1], ",")
			switch n {
			case "broadcaster":
				ms[n] = &broadcaster{n, parts}
			default:
				switch n[0] {
				case '%':
					ms[n[1:]] = &flipFlop{n[1:], Closed, parts}
				case '&':
					c := &conjuction{n[1:], map[string]pulseKind{}, parts}
					ms[n[1:]] = c
					cs = append(cs, c)
				}
			}
		}
		for _, c := range cs {
			for _, m := range ms {
				if slices.Contains(m.cs(), c.N) {
					c.Mem[m.n()] = Low
				}
			}
		}
	}

	low, high, p := 0, 0, 1

	psNext := []pulse{}
	addPulses := func(p ...pulse) {
		for _, p := range p {
			if p.K == Low {
				low++
			} else {
				high++
			}
		}
		psNext = append(psNext, p...)
	}

buttonPressLoop:
	for ; p <= maxPresses; p++ {
		addPulses(pulse{"button", "broadcaster", Low})
		for len(psNext) > 0 {
			for len(psNext) > 0 {
				ps := psNext
				psNext = []pulse{}
				for _, p := range ps {
					if stop(p) {
						break buttonPressLoop
					}
					t, found := ms[p.T]
					if !found {
						continue
					}

					if out := t.input(p); out != nil {
						addPulses(out...)
					}
				}
			}

			for _, m := range ms {
				if out := m.cycle(); out != nil {
					addPulses(out...)
				}
			}
		}
	}

	return low, high, p
}

func mapTo[S any, T any](elems []S, f func(S) T) []T {
	mapped := []T{}
	for _, e := range elems {
		mapped = append(mapped, f(e))
	}
	return mapped
}

type module interface {
	input(p pulse) []pulse
	cycle() []pulse
	cs() []string
	n() string
}

type pulseKind int8

const (
	Low  = pulseKind(1)
	High = pulseKind(2)
)

type pulse struct {
	S string
	T string
	K pulseKind
}

type broadcaster struct {
	N  string
	Cs []string
}

func (b *broadcaster) cs() []string {
	return b.Cs
}

func (b *broadcaster) cycle() []pulse {
	return nil
}

func (b *broadcaster) input(p pulse) []pulse {
	return mapTo(b.Cs, func(c string) pulse {
		return pulse{b.N, c, Low}
	})
}

func (b *broadcaster) n() string {
	return b.N
}

type conjuction struct {
	N   string
	Mem map[string]pulseKind
	Cs  []string
}

func (c *conjuction) cs() []string {
	return c.Cs
}

func (c *conjuction) cycle() []pulse {
	return nil
}

func (c *conjuction) input(p pulse) []pulse {
	c.Mem[p.S] = p.K
	out := Low
	for _, k := range c.Mem {
		if k == Low {
			out = High
		}
	}
	return mapTo(c.Cs, func(conn string) pulse {
		return pulse{c.N, conn, out}
	})
}

func (c *conjuction) n() string {
	return c.N
}

type flipFlopState int8

const (
	Closed  = flipFlopState(1)
	Opening = flipFlopState(2)
	Open    = flipFlopState(4)
	Closing = flipFlopState(8)
)

type flipFlop struct {
	N  string
	S  flipFlopState
	Cs []string
}

func (f *flipFlop) cs() []string {
	return f.Cs
}

func (f *flipFlop) cycle() []pulse {
	switch f.S {
	case Opening:
		f.S = Open
		return mapTo(f.Cs, func(c string) pulse {
			return pulse{f.N, c, High}
		})
	case Closing:
		f.S = Closed
		return mapTo(f.Cs, func(c string) pulse {
			return pulse{f.N, c, Low}
		})
	}
	return nil
}

func (f *flipFlop) input(p pulse) []pulse {
	if p.K == Low {
		switch f.S {
		case Open:
			f.S = Closing
		case Closed:
			f.S = Opening
		default:
			fmt.Fprintln(os.Stderr, "Flip flop in bad state:", f.S)
		}
	}
	return nil
}

func (f *flipFlop) n() string {
	return f.N
}
