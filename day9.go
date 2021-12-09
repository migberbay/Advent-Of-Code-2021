package main

import (
	"strconv"
	"strings"
)

func day9part1() string {
	input := string(ReadInput(9))
	h := strings.Split(input, "\n")

	heights := StringSliceToBordered2DArray(h, "9")

	risksum := 0

	var top, left, right, bottom int

	for i := 1; i < len(heights)-1; i++ {
		for j := 1; j < len(heights[i])-1; j++ {
			point, _ := strconv.Atoi(heights[i][j])

			top, _ = strconv.Atoi(heights[i-1][j])
			left, _ = strconv.Atoi(heights[i][j-1])
			right, _ = strconv.Atoi(heights[i][j+1])
			bottom, _ = strconv.Atoi(heights[i+1][j])

			if point < top && point < left && point < right && point < bottom {
				risksum += (point + 1)
			}
		}
	}

	return strconv.Itoa(risksum)
}
