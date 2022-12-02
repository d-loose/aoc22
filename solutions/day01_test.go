package solutions

import (
	"aoc2022/utils"
	"testing"
)

func TestDay01(t *testing.T) {
	testCases := []utils.TestCase{
		{Name: "example1", Input: "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", Part2: false, Output: "24000"},
		{Name: "example2", Input: "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", Part2: true, Output: "45000"},
	}
	utils.RunTests(t, Day01{}, testCases)
}
