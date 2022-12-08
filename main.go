package main

import (
	"aoc2022/solutions"
	"aoc2022/utils"
	"fmt"
)

func main() {
	days := []utils.Day{solutions.Day01{}, solutions.Day02{}, solutions.Day03{}, solutions.Day04{}, solutions.Day05{}, solutions.Day06{}, solutions.Day07{}, solutions.Day08{}}
	for _, day := range days {
		input, err := utils.GetInput(day.Name())

		if err == nil {
			fmt.Println(day.Name())
			fmt.Printf("Part 1: %s\n", day.Solution(input, false))
			fmt.Printf("Part 2: %s\n", day.Solution(input, true))
		}
	}
}
