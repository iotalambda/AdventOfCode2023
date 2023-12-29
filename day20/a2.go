package day20

import (
	"math"
	"strconv"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	/*
		Reverse engineering of the input data was needed.
		https://dreampuf.github.io/GraphvizOnline/#digraph%20G%20%7B%0Ahs%20-%3E%20sl%0Adg%20-%3E%20rx%0Avp%20-%3E%20fd%2C%20dv%0Akz%20-%3E%20jc%2C%20mc%0Anv%20-%3E%20dv%0Ahx%20-%3E%20gf%0Amm%20-%3E%20vh%0Afd%20-%3E%20td%0Adv%20-%3E%20hx%2C%20bl%2C%20rc%2C%20fd%2C%20xt%0Ahg%20-%3E%20xq%0Atd%20-%3E%20dv%2C%20hx%0Abl%20-%3E%20jt%0Abr%20-%3E%20jq%0Aqh%20-%3E%20ln%0Axq%20-%3E%20zl%2C%20cx%2C%20qh%2C%20hs%2C%20nt%2C%20sp%0Asg%20-%3E%20vv%2C%20tr%0Adm%20-%3E%20bl%2C%20dv%0Agt%20-%3E%20xq%2C%20hg%0Aln%20-%3E%20mq%2C%20xq%0Amc%20-%3E%20xv%2C%20jc%0Atx%20-%3E%20rv%2C%20jc%0Alk%20-%3E%20dg%0Amg%20-%3E%20hl%2C%20jc%0Avv%20-%3E%20zv%2C%20br%2C%20kx%2C%20mm%2C%20tr%0Ant%20-%3E%20xq%2C%20cx%0Azv%20-%3E%20dg%0Acd%20-%3E%20jc%2C%20ps%0Arc%20-%3E%20rm%2C%20dv%0Anj%20-%3E%20pt%2C%20xq%0Abroadcaster%20-%3E%20nt%2C%20kx%2C%20rc%2C%20mg%0Agf%20-%3E%20dc%2C%20dv%0Arm%20-%3E%20dm%2C%20dv%0Axx%20-%3E%20vv%2C%20cz%0Ajt%20-%3E%20dv%2C%20vp%0Azl%20-%3E%20nj%0Asp%20-%3E%20dg%0Axc%20-%3E%20jc%2C%20kz%0Axt%20-%3E%20dg%0Atp%20-%3E%20jc%0Alc%20-%3E%20vv%2C%20vn%0Avh%20-%3E%20xx%2C%20vv%0Amq%20-%3E%20hs%2C%20xq%0Acc%20-%3E%20vv%0Avn%20-%3E%20vv%2C%20cc%0Atr%20-%3E%20br%0Ahl%20-%3E%20qb%2C%20jc%0Adc%20-%3E%20dv%2C%20nv%0Ajq%20-%3E%20mm%2C%20vv%0Akx%20-%3E%20vv%2C%20sg%0Acx%20-%3E%20qh%0Asl%20-%3E%20zl%2C%20xq%0Acz%20-%3E%20lc%2C%20vv%0Aqb%20-%3E%20jc%2C%20cd%0Ajc%20-%3E%20ps%2C%20xv%2C%20lk%2C%20mg%0Axv%20-%3E%20tx%0Apt%20-%3E%20xq%2C%20gt%0Arv%20-%3E%20jc%2C%20tp%0Aps%20-%3E%20xc%0A%7D

		There are four modules `xt`, `sp`, `zv`, `lk` connected to `dg`, which is a conjuction module connected to `rx`.
		For each of the four modules there are their own dedicated "counters" created with the available module types.
		As the button is pressed, these counters cycle over the same values over and over and finally a high pulse is sent to `dg` at the end of the cycle.
		We need to find the number of presses after which all the counters send a high pulse at the same time, i.e. we need to find the LCM of the cycles.
	*/
	xt := getPressesUntil("xt", "dg", High)
	sp := getPressesUntil("sp", "dg", High)
	zv := getPressesUntil("zv", "dg", High)
	lk := getPressesUntil("lk", "dg", High)
	result := lcm(xt, sp, zv, lk)
	clipboard.WriteAll(strconv.Itoa(result))
}

func getPressesUntil(s string, t string, k pulseKind) int {
	_, _, presses := runAssignment1("day20/input.txt", math.MaxInt, func(p pulse) bool {
		return p.S == s && p.T == t && p.K == k
	})

	return presses
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, ints ...int) int {
	result := a * b / gcd(a, b)
	for _, i := range ints {
		result = lcm(result, i)
	}
	return result
}
