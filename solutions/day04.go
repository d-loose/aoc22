package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day04 struct{}

func (Day04) Name() string {
	return "day04"
}

type interval struct {
	start int
	end   int
}

func parseInterval(input string) interval {
	boundaries := strings.Split(input, "-")
	start, _ := strconv.Atoi(boundaries[0])
	end, _ := strconv.Atoi(boundaries[1])
	return interval{start, end}
}

func (i interval) containsNum(num int) bool {
	return i.start <= num && i.end >= num
}

func (i interval) contains(other interval) bool {
	return i.containsNum(other.start) && i.containsNum(other.end)
}

func (i interval) overlaps(other interval) bool {
	return i.containsNum(other.start) || i.containsNum(other.end)

}

func (Day04) Solution(input string, part2 bool) string {
	total := 0
	for _, pairs := range strings.Split(input, "\n") {
		sections := strings.Split(pairs, ",")
		interval1 := parseInterval(sections[0])
		interval2 := parseInterval(sections[1])
		if part2 && interval1.overlaps(interval2) {
			total++
		} else if interval1.contains(interval2) || interval2.contains(interval1) {
			total++
		}
	}
	return fmt.Sprint(total)
}
