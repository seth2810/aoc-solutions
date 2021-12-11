package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

// @see https://adventofcode.com/2021/day/7

const comma = ','

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

	var crabs []float64
	min, max := math.MaxFloat64, float64(0)

	for scanner.Scan() {
		f, _ := strconv.ParseFloat(scanner.Text(), 64)

		if f > max {
			max = f
		}

		if f < min {
			min = f
		}

		crabs = append(crabs, f)
	}

	minFuel := math.MaxFloat64

	for p := min; p <= max; p += 1 {
		var total float64

		for _, v := range crabs {
			total += math.Abs(v - p)
		}

		if total < minFuel {
			minFuel = total
		}
	}

	fmt.Fprintln(os.Stdout, minFuel)
}
