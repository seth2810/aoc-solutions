package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// @see https://adventofcode.com/2021/day/4

const (
	size  int    = 5
	comma string = ","
)

type board struct {
	score int
	marks [2 * size]int
	cells [size][size]int
}

func readGameBoard(scanner *bufio.Scanner) *board {
	b := &board{}

	for i, r := range b.cells {
		if ok := scanner.Scan(); !ok {
			return nil
		}

		numsScanner := bufio.NewScanner(bytes.NewReader(scanner.Bytes()))
		numsScanner.Split(bufio.ScanWords)

		for j := range r {
			if ok := numsScanner.Scan(); !ok {
				return nil
			}

			n, _ := strconv.Atoi(numsScanner.Text())

			b.score += n
			b.cells[i][j] = n
		}
	}

	return b
}

func markBoard(b *board, v int) {
	for i, r := range b.cells {
		for j, c := range r {
			if c != v {
				continue
			}

			b.score -= c
			b.marks[i] += 1
			b.marks[size+j] += 1
		}
	}
}

func isWinner(b *board) bool {
	for n := 0; n < size; n += 1 {
		if b.marks[n] == size || b.marks[size+n] == size {
			return true
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if ok := scanner.Scan(); !ok {
		return
	}

	var boards []*board
	steps := strings.Split(scanner.Text(), comma)

	for scanner.Scan() {
		boards = append(boards, readGameBoard(scanner))
	}

	for _, s := range steps {
		v, _ := strconv.Atoi(s)

		for _, b := range boards {
			markBoard(b, v)

			if isWinner(b) {
				fmt.Fprintln(os.Stdout, b.score*v)
				return
			}
		}
	}
}
