package day24

import (
	"math"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	hailstones := read("day24/input.txt")
	var h1, h2 trajectory
	{
		dist := 0.
		for haIx, ha := range hailstones {
			for _, hb := range hailstones[0:haIx] {
				if currDist := ha.dist(hb); currDist > dist {
					dist, h1, h2 = currDist, ha, hb
				}
			}
		}
	}

	t1, nextT1, t2, nextT2, err := 0., 0., 0., 0., math.MaxFloat64

	calcErr := func(t1, t2 float64) float64 {
		c1, c2 := h1.posAt(t1), h2.posAt(t2)
		rock := trajectory{c1, c2.subtract(c1)}
		err := 0.
		for _, h := range hailstones {
			err += rock.dist(h)
		}
		return err
	}

	tryNeighborhoodMove := func(s1, s2 float64) bool {
		if myT1, myT2 := t1+s1, t2+s2; myT1 >= 0 && myT2 >= 0 {
			if myErr := calcErr(myT1, myT2); myErr < err {
				nextT1, nextT2, err = myT1, myT2, myErr
				return true
			}
		}
		return false
	}

	tryAllNeighborhoodMoves := func(s float64) bool {
		a := false
		a = tryNeighborhoodMove(s, 0) || a
		a = a || tryNeighborhoodMove(-s, 0)
		a = tryNeighborhoodMove(0, s) || a
		a = a || tryNeighborhoodMove(0, -s)
		return a
	}

	for e := 10; e >= -5; e-- {
		for tryAllNeighborhoodMoves(math.Pow10(e)) {
			t1, t2 = nextT1, nextT2
			continue
		}
	}

	c1, c2 := h1.posAt(t1), h2.posAt(t2)
	fromC1ToC2 := c2.subtract(c1)
	dC1C2 := fromC1ToC2.length()
	tC1C2 := t2 - t1
	vRock := dC1C2 / tC1C2
	dRockP0C1 := vRock * t1
	rockUnit := fromC1ToC2.unit()
	rockP0 := c1.subtract(rockUnit.multiply(dRockP0C1))

	clipboard.WriteAll(strconv.Itoa((int)(rockP0.z + rockP0.y + rockP0.x)))
}

func (v v3d) length() float64 {
	return math.Sqrt(v.z*v.z + v.y*v.y + v.x*v.x)
}

func (v v3d) unit() v3d {
	len := v.length()
	return v3d{v.z / len, v.y / len, v.x / len}
}

func (tr trajectory) posAt(t float64) v3d {
	return v3d{
		tr.p0.z + tr.v.z*t,
		tr.p0.y + tr.v.y*t,
		tr.p0.x + tr.v.x*t,
	}
}

func (v v3d) multiply(f float64) v3d {
	return v3d{v.z * f, v.y * f, v.x * f}
}

func (v1 v3d) subtract(v2 v3d) v3d {
	return v3d{
		v1.z - v2.z,
		v1.y - v2.y,
		v1.x - v2.x,
	}
}

func (v1 v3d) cross(v2 v3d) v3d {
	return v3d{
		v1.x*v2.y - v1.y*v2.x,
		v1.z*v2.x - v1.x*v2.z,
		v1.y*v2.z - v1.z*v2.y,
	}
}

func (v1 v3d) dot(v2 v3d) float64 {
	return v1.z*v2.z + v1.y*v2.y + v1.x*v2.x
}

func (tr1 trajectory) dist(tr2 trajectory) float64 {
	c := tr1.v.cross(tr2.v)
	return math.Abs(c.dot(tr1.p0.subtract(tr2.p0))) / c.length()
}
