package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// @see https://adventofcode.com/2021/day/1

func main() {
	var p int
	var n int
	var c int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())

		if p == 0 {
			p = n
			continue
		}

		if n-p > 0 {
			c += 1
		}

		p = n
	}

	fmt.Fprintln(os.Stdout, c)
}
