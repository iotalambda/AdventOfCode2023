package day25

import (
	"aoc2023/utils"
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unsafe"

	"github.com/atotto/clipboard"
)

func Assignment1() {
	file, err := os.Open("day25/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g := newGraph()
	for scanner.Scan() {
		ps := strings.Split(scanner.Text(), " ")
		srcVK := strings.Replace(ps[0], ":", "", 1)
		tgtVKs := ps[1:]
		g.ensureConnections(srcVK, tgtVKs)
	}

	var lastRemoved *edge
	var sampleCrossSgPairsOrNil map[vPairKey]struct{}

	for r := 0; r < 3; r++ {

		edgeHistCache := map[vPairKey][]*edge{}
		vertices, edges := utils.Array(g.vertices), utils.Array(g.edges)

		for vIx1, v1 := range vertices {
			for _, v2 := range vertices[0:vIx1] {

				key := newVPairKey(v1, v2)
				if sampleCrossSgPairsOrNil != nil {
					if _, ok := sampleCrossSgPairsOrNil[key]; !ok {
						continue
					}
				}

				var shPthEdges []*edge
				if cached, ok := edgeHistCache[key]; ok {
					shPthEdges = cached
				} else {
					v1.visit, v2.visit = newVisit(v1), newVisit(v2)
					nextVs := []*vertex{v1}

				shPthSearch:
					for {
						currVs := nextVs
						nextVs = []*vertex{}
						for _, v1 := range currVs {
							for _, e := range v1.edges {
								v2 := e.getOther(v1)
								if v2.visit == nil {
									v2.visit = v1.visit.with(e)
									nextVs = append(nextVs, v2)

									vO := v2.visit.origin
									for hEIx, hE := range v2.visit.edgeHist {
										edgeHistCache[newVPairKey(vO, v2)] = v2.visit.edgeHist[hEIx:]
										vO = hE.getOther(vO)
										if vO == v2 {
											break
										}
									}
								} else if v2.visit.origin != v1.visit.origin {
									shPthEdges = make([]*edge, 0, len(v2.visit.edgeHist)+len(v1.visit.edgeHist)+1)
									shPthEdges = append(shPthEdges, v2.visit.edgeHist...)
									shPthEdges = append(shPthEdges, v1.visit.edgeHist...)
									shPthEdges = append(shPthEdges, e)
									break shPthSearch
								}
							}
						}
					}
				}
				for _, e := range shPthEdges {
					e.totShPths = append(e.totShPths, key)
				}
				g.resetVisit()
			}
		}

		lastRemoved = slices.MaxFunc(edges, func(e1, e2 *edge) int {
			return cmp.Compare(len(e1.totShPths), len(e2.totShPths))
		})
		g.removeEdge(lastRemoved)

		if r == 0 {
			sampleCrossSgPairsOrNil = map[vPairKey]struct{}{}
			for _, k := range lastRemoved.totShPths[0:min(100, len(lastRemoved.totShPths))] {
				sampleCrossSgPairsOrNil[k] = struct{}{}
			}
		}
		g.resetTot()
	}

	sgSize := func(v *vertex) int {
		size, nextVs := 0, []*vertex{v}
		v.visit = newVisit(v)
		for len(nextVs) > 0 {
			size += len(nextVs)
			currVs := nextVs
			nextVs = []*vertex{}
			for _, v1 := range currVs {
				for _, e := range v1.edges {
					if v2 := e.getOther(v1); v2.visit == nil {
						v2.visit = newVisit(v)
						nextVs = append(nextVs, v2)
					}
				}
			}
		}
		g.resetVisit()
		return size
	}

	clipboard.WriteAll(strconv.Itoa(sgSize(lastRemoved.v1) * sgSize(lastRemoved.v2)))
}

type graph struct {
	edges    map[vPairKey]*edge
	vertices map[string]*vertex
}

type edge struct {
	key       vPairKey
	v1, v2    *vertex
	totShPths []vPairKey
}

type vertex struct {
	key   string
	edges []*edge
	visit *visit
}

type visit struct {
	origin   *vertex
	edgeHist []*edge
}

type vPairKey struct {
	va, vb *vertex
}

func newGraph() graph {
	return graph{map[vPairKey]*edge{}, map[string]*vertex{}}
}

func (g *graph) ensureConnections(srcVK string, tgtVKs []string) {
	srcV := getOrAddVertex(g, srcVK)
	for _, tgtVK := range tgtVKs {
		tgtV := getOrAddVertex(g, tgtVK)
		getOrAddEdge(g, srcV, tgtV)
	}
}

func (g *graph) removeEdge(e *edge) {
	delete(g.edges, e.key)
	slices.DeleteFunc(g.vertices[e.v1.key].edges, func(e2 *edge) bool {
		return e2 == e
	})
	slices.DeleteFunc(g.vertices[e.v2.key].edges, func(e2 *edge) bool {
		return e2 == e
	})
}

func (g *graph) resetVisit() {
	for _, v := range g.vertices {
		v.visit = nil
	}
}

func (g *graph) resetTot() {
	for _, e := range g.edges {
		e.totShPths = []vPairKey{}
	}
}

func getOrAddVertex(g *graph, key string) *vertex {
	v, ok := g.vertices[key]
	if !ok {
		v = &vertex{key, []*edge{}, nil}
		g.vertices[key] = v
	}
	return v
}

func getOrAddEdge(g *graph, v1, v2 *vertex) *edge {
	key := newVPairKey(v1, v2)
	e, ok := g.edges[key]
	if !ok {
		e = &edge{key, v1, v2, []vPairKey{}}
		g.edges[key] = e
		v1.edges = append(v1.edges, e)
		v2.edges = append(v2.edges, e)
	}
	return e
}

func newVPairKey(v1, v2 *vertex) vPairKey {
	if uintptr(unsafe.Pointer(v2)) < uintptr(unsafe.Pointer(v1)) {
		v1, v2 = v2, v1
	}
	return vPairKey{v1, v2}
}

func newVisit(origin *vertex) *visit {
	return &visit{origin, []*edge{}}
}

func (p *visit) with(e *edge) *visit {
	return &visit{p.origin, append(slices.Clone(p.edgeHist), e)}
}

func (e *edge) getOther(v *vertex) *vertex {
	if e.v1 == v {
		return e.v2
	}

	if e.v2 == v {
		return e.v1
	}

	fmt.Fprintln(os.Stderr, "Vertex not found for the edge.")
	os.Exit(1)
	panic(nil)
}
