package day22

import (
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	runAssignment1("day22/input.txt", func(w world, bs []*brick) {
		falling := 0

		for ixBR, bR := range bs {
			w.removeBrick(bR)
			bsO := make([]*brick, 0)
			bsD := make([]*brick, 0)
			for ixBD, bD := range bs {
				if ixBR >= ixBD {
					continue
				}
				if d := w.getDropDistance(bD); d > 0 {
					w.removeBrick(bD)
					c := *bD
					bD2 := &c
					w.dropBrick(bD2, d)
					w.addBrick(bD2)
					bsO = append(bsO, bD)
					bsD = append(bsD, bD2)
				}
			}
			falling += len(bsD)
			for _, bD := range bsD {
				w.removeBrick(bD)
			}
			for _, bO := range bsO {
				w.addBrick(bO)
			}
			w.addBrick(bR)
		}

		clipboard.WriteAll(strconv.Itoa(falling))
	})
}
