package main

import (
	"strconv"
	"strings"
)

func day1part1() string {
	input := string(ReadInput(1))

	prevDepth := 2000
	increases := 0

	for _, line := range strings.Split(input, "\n") {
		currDepth, _ := strconv.Atoi(strings.TrimSpace(line))
		if currDepth > prevDepth {
			increases++
		}
		prevDepth = currDepth
	}

	return strconv.Itoa(increases)
}

func day1part2() string {
	input := string(ReadInput(1))

	prevDepth := 2000
	increases := 0

	lines := strings.Split(input, "\n")

	for i := 1; i < len(lines)-1; i++ {
		currD1, _ := strconv.Atoi(strings.TrimSpace(lines[i-1]))
		currD2, _ := strconv.Atoi(strings.TrimSpace(lines[i]))
		currD3, _ := strconv.Atoi(strings.TrimSpace(lines[i+1]))
		currDepth := currD1 + currD2 + currD3

		if currDepth > prevDepth {
			increases++
		}
		prevDepth = currDepth
	}

	return strconv.Itoa(increases)
}
