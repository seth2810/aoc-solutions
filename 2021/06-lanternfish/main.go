package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// @see https://adventofcode.com/2021/day/6

const (
	comma            = ','
	days             = 80
	fishInitialTimer = 8
	fishResetTimer   = 6
)

func splitByComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	commaIndex := bytes.IndexByte(data, comma)
	if commaIndex > 0 {
		buffer := data[:commaIndex]
		return commaIndex + 1, bytes.TrimSpace(buffer), nil
	}

	if atEOF && len(data) != 0 {
		return len(data), bytes.TrimSpace(data), nil
	}

	return 0, nil, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(splitByComma)

	var fishes []int

	for scanner.Scan() {
		t, _ := strconv.Atoi(scanner.Text())

		fishes = append(fishes, t)
	}

	for i := 1; i <= days; i += 1 {
		rest := make([]int, 0)

		for j, v := range fishes {
			if v == 0 {
				fishes[j] = fishResetTimer
				rest = append(rest, fishInitialTimer)
				continue
			}

			fishes[j] = v - 1
		}

		fishes = append(fishes, rest...)
	}

	fmt.Fprintln(os.Stdout, len(fishes))
}
