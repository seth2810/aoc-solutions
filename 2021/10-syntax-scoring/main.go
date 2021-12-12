package main

import (
	"bufio"
	"fmt"
	"os"
)

// @see https://adventofcode.com/2021/day/10

var brackets = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var points = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sum int

outer:
	for scanner.Scan() {
		stack := make([]byte, 0, len(scanner.Bytes()))

		for _, v := range scanner.Bytes() {
			if _, ok := brackets[v]; ok {
				stack = append(stack, v)
				continue
			}

			if _, ok := points[v]; !ok {
				continue
			}

			if c := brackets[stack[len(stack)-1]]; v != c {
				if p, ok := points[v]; ok {
					sum += p
				}

				continue outer
			}

			stack = stack[:len(stack)-1]
		}

	}

	fmt.Fprintln(os.Stdout, sum)
}
