package solutions

import (
	"aoc2022/utils"
	"testing"
)

func TestDay04(t *testing.T) {
	testCases := []utils.TestCase{
		{Name: "example1", Input: "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8", Part2: false, Output: "2"},
		{Name: "example1", Input: "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8", Part2: true, Output: "4"},
	}
	utils.RunTests(t, Day04{}, testCases)
}
