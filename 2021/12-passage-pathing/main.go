package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// @see https://adventofcode.com/2021/day/12

const (
	start cave = "start"
	end   cave = "end"
)

type cave string
type route []cave
type caves map[cave][]cave

func isSmall(c cave) bool {
	return c[0]-97 <= 25
}

func countRoutes(cs caves, c cave, visited route) int {
	if c == end {
		return 1
	}

	smallsVisited := make(map[cave]struct{})
	for _, v := range visited {
		if isSmall(v) {
			smallsVisited[v] = struct{}{}
		}
	}

	if _, ok := smallsVisited[c]; ok {
		return 0
	}

	neighbours := cs[c]
	nexts := make([]cave, 0, len(neighbours))
	for _, n := range neighbours {
		if _, ok := smallsVisited[n]; !ok {
			nexts = append(nexts, n)
		}
	}

	var total int
	for _, n := range nexts {
		total += countRoutes(cs, n, append(route{c}, visited...))
	}

	return total
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cs := make(caves)

	for scanner.Scan() {
		names := strings.SplitN(scanner.Text(), "-", 2)
		left, right := cave(names[0]), cave(names[1])

		cs[left] = append(cs[left], right)
		cs[right] = append(cs[right], left)
	}

	fmt.Fprintln(os.Stdout, countRoutes(cs, start, nil))
}
