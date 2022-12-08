package solutions

import (
	"aoc2022/utils"
	"testing"
)

func TestDay08(t *testing.T) {
	testCases := []utils.TestCase{
		{Name: "example1", Input: "30373\n25512\n65332\n33549\n35390", Part2: false, Output: "21"},
		{Name: "example2", Input: "30373\n25512\n65332\n33549\n35390", Part2: true, Output: "8"},
	}
	utils.RunTests(t, Day08{}, testCases)
}
