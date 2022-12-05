package solutions

import (
	"aoc2022/utils"
	"testing"
)

func TestDay05(t *testing.T) {
	testCases := []utils.TestCase{
		{Name: "example1", Input: "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2", Part2: false, Output: "CMZ"},
		{Name: "example2", Input: "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2", Part2: true, Output: "MCD"},
	}
	utils.RunTests(t, Day05{}, testCases)
}
