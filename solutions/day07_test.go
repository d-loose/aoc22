package solutions

import (
	"aoc2022/utils"
	"testing"
)

func TestDay07(t *testing.T) {
	testCases := []utils.TestCase{
		{Name: "example1", Input: "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k", Part2: false, Output: "95437"},
		{Name: "example2", Input: "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k", Part2: true, Output: "24933642"},
	}
	utils.RunTests(t, Day07{}, testCases)
}
