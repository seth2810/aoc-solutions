package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// @see https://adventofcode.com/2021/day/13

const (
	foldXPrefix string    = "fold along x="
	foldYPrefix string    = "fold along y="
	horizontal  direction = "horizontal"
	vertical    direction = "vertical"
)

type direction string

type fold struct {
	dir direction
	pos int
}

type field struct {
	n, m  int
	cells [][]int
}

type point struct {
	x, y int
}

func buildField(ps []*point) *field {
	var n, m int
	for _, p := range ps {
		if p.x > n {
			n = p.x
		}

		if p.y > m {
			m = p.y
		}
	}

	cells := make([][]int, m+1)
	for i := range cells {
		cells[i] = make([]int, n+1)
	}

	for _, p := range ps {
		cells[p.y][p.x] = 1
	}

	return &field{n, m, cells}
}

func foldField(f *field, fld *fold) *field {
	n, m := f.n, f.m

	switch fld.dir {
	case vertical:
		m = fld.pos
	case horizontal:
		n = fld.pos
	}

	nextCells := make([][]int, m+1)
	for i := range nextCells {
		nextCells[i] = make([]int, n+1)
	}

	for i, r := range nextCells {
		for j := range r {
			if f.cells[i][j] == 1 {
				nextCells[i][j] = 1
			}

			switch fld.dir {
			case vertical:
				if f.cells[f.m-i][j] == 1 {
					nextCells[i][j] = 1
				}
			case horizontal:
				if f.cells[i][f.n-j] == 1 {
					nextCells[i][j] = 1
				}
			}

		}
	}

	return &field{n, m, nextCells}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	folds := make([]*fold, 0)
	points := make([]*point, 0)

	for scanner.Scan() && len(scanner.Bytes()) != 0 {
		cs := strings.SplitN(scanner.Text(), ",", 2)
		x, _ := strconv.Atoi(cs[0])
		y, _ := strconv.Atoi(cs[1])

		points = append(points, &point{x, y})
	}

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), foldYPrefix) {
			p, _ := strconv.Atoi(strings.TrimPrefix(scanner.Text(), foldYPrefix))
			folds = append(folds, &fold{vertical, p})
			break
		}

		if strings.HasPrefix(scanner.Text(), foldXPrefix) {
			p, _ := strconv.Atoi(strings.TrimPrefix(scanner.Text(), foldXPrefix))
			folds = append(folds, &fold{horizontal, p})
			break
		}
	}

	field := buildField(points)
	for _, f := range folds {
		field = foldField(field, f)
	}

	var dots int
	for _, r := range field.cells {
		for _, c := range r {
			if c == 1 {
				dots += 1
			}
		}
	}

	fmt.Fprintln(os.Stdout, dots)
}
