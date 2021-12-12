package main

import (
	"bufio"
	"fmt"
	"os"
)

// @see https://adventofcode.com/2021/day/9

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var field [][]int

	for scanner.Scan() {
		row := make([]int, len(scanner.Bytes()))

		for i, v := range scanner.Bytes() {
			row[i] = int(v - 48)
		}

		field = append(field, row)
	}

	var sum int

	for i, r := range field {
	inner:
		for j, c := range r {
			if i > 0 && c >= field[i-1][j] {
				continue inner
			}

			if i < len(field)-1 && c >= field[i+1][j] {
				continue inner
			}

			if j > 0 && c >= field[i][j-1] {
				continue inner
			}

			if j < len(r)-1 && c >= field[i][j+1] {
				continue inner
			}

			sum += c + 1
		}
	}

	fmt.Fprintln(os.Stdout, sum)
}
