package main

import (
	"fmt"
)

func main() {
	hasquit := false
	for !hasquit {
		fmt.Println("What day do you want to run? (01-25), (please use 2 digits), type \"quit\" to end.")
		var day string
		fmt.Scanln(&day)

		switch day {
		case "01":
			res1 := day1part1()
			res2 := day1part2()
			PrintSolutions(res1, res2)

		case "09":
			res1 := day9part1()
			res2 := " none yet "
			PrintSolutions(res1, res2)

		case "quit":
			hasquit = true
			fmt.Println("GoodBye!")
		default:
			fmt.Errorf("That is not a valid day.")

		}
	}
}
