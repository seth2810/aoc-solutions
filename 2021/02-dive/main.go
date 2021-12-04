package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// @see https://adventofcode.com/2021/day/2

func main() {
	var v int
	var h int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for {
		if ok := scanner.Scan(); !ok {
			break
		}

		cmd := scanner.Text()

		if ok := scanner.Scan(); !ok {
			break
		}

		units, _ := strconv.Atoi(scanner.Text())

		switch cmd {
		case "forward":
			h += units
		case "up":
			v -= units
		case "down":
			v += units
		}
	}

	fmt.Fprintln(os.Stdout, h*v)
}
