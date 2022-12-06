package solutions

import (
	"aoc2022/utils"
	"testing"
)

func TestDay06(t *testing.T) {
	testCases := []utils.TestCase{
		{Name: "example1", Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Part2: false, Output: "5"},
		{Name: "example2", Input: "nppdvjthqldpwncqszvftbrmjlhg", Part2: false, Output: "6"},
		{Name: "example3", Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Part2: false, Output: "10"},
		{Name: "example4", Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Part2: false, Output: "11"},
		{Name: "example5", Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", Part2: true, Output: "19"},
		{Name: "example6", Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Part2: true, Output: "23"},
		{Name: "example7", Input: "nppdvjthqldpwncqszvftbrmjlhg", Part2: true, Output: "23"},
		{Name: "example8", Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Part2: true, Output: "29"},
		{Name: "example9", Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Part2: true, Output: "26"},
	}
	utils.RunTests(t, Day06{}, testCases)
}
