package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// @see https://adventofcode.com/2021/day/14

type key struct {
	left, right byte
}

type rules map[key][]byte

func apply(b []byte, rs rules) []byte {
	r := make([]byte, 0, len(b))

	for i, c := range b {
		if i == 0 {
			r = append(r, c)
			continue
		}

		if t, ok := rs[key{b[i-1], c}]; ok {
			r = append(r, t...)
			continue
		}

		r = append(r, c)
	}

	return r
}

func findMinMaxCounter(b []byte) (int, int) {
	acc := make(map[byte]int)
	min, max := math.MaxInt32, 0

	for _, c := range b {
		acc[c] += 1
	}

	for _, v := range acc {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return min, max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	rs := make(rules)
	t := scanner.Bytes()

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		buff := scanner.Bytes()

		rs[key{buff[0], buff[1]}] = []byte{buff[6], buff[1]}
	}

	for i := 0; i < 10; i += 1 {
		t = apply(t, rs)
	}

	min, max := findMinMaxCounter(t)

	fmt.Fprintln(os.Stdout, max-min)
}
