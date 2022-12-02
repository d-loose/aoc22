package main

import (
	"aoc2022/solutions"
	"fmt"
	"os"
	"path"
	"strings"
)

type Day interface {
	Name() string
	Solution(string, bool) string
}

func getInput(name string) (string, error) {
	input, err := os.ReadFile(path.Join("inputs", name))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(input)), nil
}

func main() {
	days := []Day{solutions.Day01{}, solutions.Day02{}}
	for _, day := range days {
		input, err := getInput(day.Name())

		if err == nil {
			fmt.Println(day.Name())
			fmt.Printf("Part 1: %s\n", day.Solution(input, false))
			fmt.Printf("Part 2: %s\n", day.Solution(input, true))
		}
	}
}
