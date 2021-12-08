package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// @see https://adventofcode.com/2021/day/5

const (
	comma     string = ","
	separator string = " -> "
)

type vent struct {
	x1, y1, x2, y2 int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var maxX int
	var maxY int
	var vents []*vent

	for scanner.Scan() {
		ends := strings.Split(scanner.Text(), separator)
		crds1 := strings.Split(ends[0], comma)
		crds2 := strings.Split(ends[1], comma)

		if crds1[0] == crds2[0] || crds1[1] == crds2[1] {
			x1, _ := strconv.ParseFloat(crds1[0], 64)
			y1, _ := strconv.ParseFloat(crds1[1], 64)
			x2, _ := strconv.ParseFloat(crds2[0], 64)
			y2, _ := strconv.ParseFloat(crds2[1], 64)

			v := &vent{
				int(math.Min(x1, x2)),
				int(math.Min(y1, y2)),
				int(math.Max(x1, x2)),
				int(math.Max(y1, y2)),
			}

			if v.x2 > maxX {
				maxX = v.x2
			}

			if v.y2 > maxY {
				maxY = v.y2
			}

			vents = append(vents, v)
		}
	}

	field := make([][]int, maxY+1)
	for i := range field {
		field[i] = make([]int, maxX+1)
	}

	for _, v := range vents {
		if v.x1 != v.x2 {
			for i := v.x1; i <= v.x2; i += 1 {
				field[v.y1][i] += 1
			}
		} else {
			for j := v.y1; j <= v.y2; j += 1 {
				field[j][v.x1] += 1
			}
		}
	}

	var count int

	for _, r := range field {
		for _, c := range r {
			if c >= 2 {
				count += 1
			}
		}
	}

	fmt.Fprintln(os.Stdout, count)
}
