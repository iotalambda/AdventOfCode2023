package day21

import (
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	dMax := 26501365
	g := load("day21/input.txt")
	w := len(*g)
	c := (w - 3) / 2
	yS, xS := g.find('S')
	(*g)[yS][xS] = '.'

	fF := (dMax - w + 1) / w
	nNA := fF / 2
	nA := 4*nNA*(nNA+1) + 1
	nNB := (fF + 1) / 2
	nB := 4 * nNB * nNB

	var sO, sS, eI int
	if (dMax-c-2)%w <= c {
		sO = (dMax - c - 2) % w
		sI := sO + w
		sS = sI - c - 2

		eI = g.runAssignment1(0, c+1, sI)
		eI += g.runAssignment1(w-1, c+1, sI)
		eI += g.runAssignment1(c+1, 0, sI)
		eI += g.runAssignment1(c+1, w-1, sI)
	} else {
		sO = (dMax - c - 2) % w
		sS = sO - c - 2
	}

	sL := sS + w
	nS := (dMax - c - 1) / w
	nL := nS - 1

	eA := g.runAssignment1(yS, xS, dMax)
	eA *= nA
	eB := g.runAssignment1(yS, xS, dMax+1)
	eB *= nB

	eO := g.runAssignment1(0, c+1, sO)
	eO += g.runAssignment1(w-1, c+1, sO)
	eO += g.runAssignment1(c+1, 0, sO)
	eO += g.runAssignment1(c+1, w-1, sO)

	eS := g.runAssignment1(0, 0, sS)
	eS += g.runAssignment1(w-1, 0, sS)
	eS += g.runAssignment1(0, w-1, sS)
	eS += g.runAssignment1(w-1, w-1, sS)
	eS *= nS

	eL := g.runAssignment1(0, 0, sL)
	eL += g.runAssignment1(w-1, 0, sL)
	eL += g.runAssignment1(0, w-1, sL)
	eL += g.runAssignment1(w-1, w-1, sL)
	eL *= nL

	result := eA + eB + eI + eO + eS + eL

	clipboard.WriteAll(strconv.Itoa(result))
}
