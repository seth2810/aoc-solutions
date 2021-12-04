package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// @see https://adventofcode.com/2021/day/3

const (
	one  byte = '1'
	zero byte = '0'
)

func main() {
	var n int
	var onesCount []int
	var zerosCount []int
	var gammaBytes []byte
	var epsilonBytes []byte
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if n == 0 {
			n = len(line)
			onesCount = make([]int, n)
			zerosCount = make([]int, n)
			gammaBytes = make([]byte, n)
			epsilonBytes = make([]byte, n)
		}

		for i := 0; i < n; i += 1 {
			switch line[i] {
			case zero:
				zerosCount[i] += 1
			case one:
				onesCount[i] += 1
			}
		}
	}

	for i := 0; i < n; i += 1 {
		if onesCount[i] > zerosCount[i] {
			gammaBytes[i] = one
			epsilonBytes[i] = zero
		} else {
			gammaBytes[i] = zero
			epsilonBytes[i] = one
		}
	}

	gamma, _ := strconv.ParseInt(string(gammaBytes), 2, 64)
	epsilon, _ := strconv.ParseInt(string(epsilonBytes), 2, 64)

	fmt.Fprintln(os.Stdout, gamma*epsilon)
}
