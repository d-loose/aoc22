package solutions

import (
	"fmt"
	"strings"
)

type Day03 struct{}

func (Day03) Name() string {
	return "day03"
}

type set map[rune]struct{}

func stringToSet(str string) set {
	var s = make(set)
	for _, r := range str {
		s[r] = struct{}{}
	}
	return s
}

func intersect(set1, set2 set) set {
	var i = make(set)
	for r := range set1 {
		if _, ok := set2[r]; ok {
			i[r] = struct{}{}
		}
	}
	return i
}

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}
	return 0
}

func (Day03) Solution(input string, part2 bool) string {
	total := 0
	var intersection set
	for i, rucksack := range strings.Split(input, "\n") {
		if part2 {
			if i%3 == 0 {
				intersection = stringToSet(rucksack)
			} else {
				intersection = intersect(intersection, stringToSet(rucksack))
			}
		} else {
			compartment1 := stringToSet(rucksack[:len(rucksack)/2])
			compartment2 := stringToSet(rucksack[len(rucksack)/2:])
			intersection = intersect(compartment1, compartment2)
		}
		if !part2 || i%3 == 2 {
			for item := range intersection {
				total += priority(item)
			}
		}
	}
	return fmt.Sprint(total)
}
