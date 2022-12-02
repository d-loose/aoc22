package main

import (
	"aoc2022/solutions"
	"aoc2022/utils"
	"fmt"
)

func main() {
	days := []utils.Day{solutions.Day01{}, solutions.Day02{}}
	for _, day := range days {
		input, err := utils.GetInput(day.Name())

		if err == nil {
			fmt.Println(day.Name())
			fmt.Printf("Part 1: %s\n", day.Solution(input, false))
			fmt.Printf("Part 2: %s\n", day.Solution(input, true))
		}
	}
}
