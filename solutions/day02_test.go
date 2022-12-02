package solutions

import (
	"aoc2022/utils"
	"testing"
)

func TestDay02(t *testing.T) {
	testCases := []utils.TestCase{
		{Name: "example1", Input: "A Y\nB X\nC Z", Part2: false, Output: "15"},
		{Name: "example2", Input: "A Y\nB X\nC Z", Part2: true, Output: "12"},
	}
	utils.RunTests(t, Day02{}, testCases)
}
