package main

import (
	"bufio"
	"fmt"
	"os"
)

// @see https://adventofcode.com/2021/day/11

type octopus struct {
	value      int
	flashes    int
	lastFlash  int
	neighbours []*octopus
}

type cave [][]*octopus

func findNeighbours(c cave) {
	var n, m int

	for i, r := range c {
		if i == 0 {
			n, m = len(c), len(r)
		}

		for j, o := range r {
			if i > 0 {
				o.neighbours = append(o.neighbours, c[i-1][j])

				if j > 0 {
					o.neighbours = append(o.neighbours, c[i-1][j-1])
				}

				if j < m-1 {
					o.neighbours = append(o.neighbours, c[i-1][j+1])
				}
			}

			if i < n-1 {
				o.neighbours = append(o.neighbours, c[i+1][j])

				if j > 0 {
					o.neighbours = append(o.neighbours, c[i+1][j-1])
				}

				if j < m-1 {
					o.neighbours = append(o.neighbours, c[i+1][j+1])
				}
			}

			if j > 0 {
				o.neighbours = append(o.neighbours, c[i][j-1])
			}

			if j < m-1 {
				o.neighbours = append(o.neighbours, c[i][j+1])
			}
		}
	}
}

func flash(s int, o *octopus) {
	if o.lastFlash == s {
		return
	}

	o.value += 1

	if o.value <= 9 {
		return
	}

	o.value = 0
	o.flashes += 1
	o.lastFlash = s

	for _, n := range o.neighbours {
		flash(s, n)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var c cave

	for scanner.Scan() {
		r := make([]*octopus, 0, len(scanner.Bytes()))

		for _, v := range scanner.Bytes() {
			r = append(r, &octopus{value: int(v - 48)})
		}

		c = append(c, r)
	}

	findNeighbours(c)

	for s := 0; s <= 100; s += 1 {
		for _, r := range c {
			for _, o := range r {
				flash(s, o)
			}
		}
	}

	var flashes int
	for _, r := range c {
		for _, o := range r {
			flashes += o.flashes
		}
	}

	fmt.Fprintln(os.Stdout, flashes)
}
