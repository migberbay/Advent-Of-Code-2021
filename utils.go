package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func ReadInput(day int) []byte {
	folder := "day" + strconv.Itoa(day)

	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	exPath := ex + "/days/" + folder + ".txt"

	body, err := ioutil.ReadFile(exPath)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return body
}

func PrintSolutions(solution1 string, solution2 string) {
	fmt.Printf("the first solution is: %s and the second is: %s\n", solution1, solution2)
}

func StringSliceToBordered2DArray(input []string, borderchar string) [][]string {
	h := len(input) + 2
	w := len(input[0]) + 2

	res := make([][]string, h)

	for r := range res {
		res[r] = make([]string, w)
	}

	for i := 0; i < h; i++ {
		fmt.Printf("line %v \n", i)
		for j := 0; j < w; j++ {
			if i == 0 || i == w-1 || j == 0 || j == h-1 {
				res[i][j] = borderchar
			} else {
				res[i][j] = string(input[i-1][j-1])
			}
		}
	}

	return res
}

// func AddBorderToDoubleArray(input [][]string, borderchar string) [][]string {
// 	w := len(input[0]) + 2
// 	h := len(input) + 2

// 	res := make([][]string, 0)

// 	for i := 0; i < h; i++ {
// 		for j := 0; j < w; j++ {
// 			if i == 0 || i == w-1 || j == 0 || j == h-1 {
// 				res[i][j] = borderchar
// 			} else {
// 				res[i][j] = input[i-1][j-1]
// 			}
// 		}
// 	}
// 	return res
// }
