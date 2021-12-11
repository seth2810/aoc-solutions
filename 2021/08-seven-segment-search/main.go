package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// @see https://adventofcode.com/2021/day/8

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var count int

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), " | ", 2)

		segScanner := bufio.NewScanner(bytes.NewBufferString(parts[1]))
		segScanner.Split(bufio.ScanWords)

		for segScanner.Scan() {
			switch len(segScanner.Bytes()) {
			case 2, 3, 4, 7:
				count += 1
			}
		}
	}

	fmt.Fprintln(os.Stdout, count)
}
