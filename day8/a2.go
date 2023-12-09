package day8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func Assignment2() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var instructions []rune
	var i_tot int
	letters := int('Z' - 'A' + 1)
	network := make([][2]int, int(math.Pow(float64(letters), 3)))
	a_ub := letters * letters
	z_lb := (letters - 1) * letters * letters
	a_nodes := make([]int, 0)

	for l_ix := 0; scanner.Scan(); l_ix++ {
		if l_ix == 0 {
			instructions = []rune(scanner.Text())
			i_tot = len(instructions)
			continue
		}

		if l_ix == 1 {
			continue
		}

		line := strings.Fields(scanner.Text())
		node := int(line[0][0]-'A') + int(line[0][1]-'A')*letters + int(line[0][2]-'A')*letters*letters
		l := int(line[2][1]-'A') + int(line[2][2]-'A')*letters + int(line[2][3]-'A')*letters*letters
		r := int(line[3][0]-'A') + int(line[3][1]-'A')*letters + int(line[3][2]-'A')*letters*letters
		network[node] = [...]int{l, r}
		if node < a_ub {
			a_nodes = append(a_nodes, node)
		}
	}

	node_n_z_steps := make([]int, len(a_nodes))
	node_cycl_rnds := make([]int, len(a_nodes))

a_nodes_loop:
	for n_ix, n := range a_nodes {
		is_ns := make([][]int, i_tot)
		z_r_ix := 0
		z_i_ix := 0
		for r_ix := 0; ; r_ix++ {
			for i_ix, i := range instructions {
				if i == 'L' {
					n = network[n][0]
				} else {
					n = network[n][1]
				}

				if r_ix == 0 {
					is_ns[i_ix] = make([]int, 0)
				}

				curr_i_ns := is_ns[i_ix]
				if c_r_ix := slices.Index(curr_i_ns, n); c_r_ix != -1 {
					node_n_z_steps[n_ix] = z_r_ix*i_tot + z_i_ix
					continue a_nodes_loop
				}

				is_ns[i_ix] = append(curr_i_ns, n)

				if n > z_lb {
					z_r_ix = r_ix
					z_i_ix = i_ix
				}
			}
		}
	}

	max_cycl_rnds_ix := slices.Index(node_cycl_rnds, slices.Max(node_cycl_rnds))

a_loop:
	for a := 0; ; a++ {
		total_steps := node_n_z_steps[max_cycl_rnds_ix] + a*i_tot*node_cycl_rnds[max_cycl_rnds_ix]
		for n_ix, cycl_rnds := range node_cycl_rnds {
			if n_ix == max_cycl_rnds_ix {
				continue
			}
			if (total_steps-node_n_z_steps[n_ix])%(i_tot*cycl_rnds) != 0 {
				continue a_loop
			}
		}
		clipboard.WriteAll(strconv.Itoa(total_steps + 1))
		break
	}
}
