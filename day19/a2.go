package day19

import (
	"maps"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	wfs, _ := loadAssignment("day19/input.txt")
	minPart, minV := part{}, 1
	minPart['x'], minPart['m'], minPart['a'], minPart['s'] = minV, minV, minV, minV
	maxPart, maxV := part{}, 4000
	maxPart['x'], maxPart['m'], maxPart['a'], maxPart['s'] = maxV, maxV, maxV, maxV

	agentsNext := []agent{{minPart, maxPart, "in"}}
	var agents []agent

	result := 0
	for len(agentsNext) > 0 {
		agents = agentsNext
		agentsNext = []agent{}
		for _, a := range agents {
		workflowLoop:
			for {
				switch a.WfN {
				case "A":
					result += a.Left('x') * a.Left('m') * a.Left('a') * a.Left('s')
					fallthrough
				case "R":
					break workflowLoop
				default:
					wf := wfs[a.WfN]
					for _, r := range wf.Rs {
						a2 := agent{maps.Clone(a.Min), maps.Clone(a.Max), a.WfN}
						if r.GT {
							a.Max[r.C] = min(a.Max[r.C], r.V)
							a2.Min[r.C] = max(a2.Min[r.C], r.V+1)
						} else {
							a.Min[r.C] = max(a.Min[r.C], r.V)
							a2.Max[r.C] = min(a2.Max[r.C], r.V-1)
						}

						if a2.Left(r.C) > 0 {
							a2.WfN = r.WfN
							agentsNext = append(agentsNext, a2)
						}

						if a.Left(r.C) <= 0 {
							break workflowLoop
						}
					}
					a.WfN = wf.F
				}
			}
		}
	}

	clipboard.WriteAll(strconv.Itoa(result))
}

type agent struct {
	Min part
	Max part
	WfN string
}

func (a *agent) Left(c category) int {
	return a.Max[c] - a.Min[c] + 1
}
