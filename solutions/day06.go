package solutions

import "fmt"

type Day06 struct{}

func (Day06) Name() string {
	return "day06"
}

func findMarker(packet string, length int) int {
	for pos := 0; pos < len(packet)-length; pos++ {
		if len(stringToSet(packet[pos:pos+length])) == length {
			return pos + length
		}
	}
	return -1
}

func (Day06) Solution(input string, part2 bool) string {
	var length int
	if part2 {
		length = 14
	} else {
		length = 4
	}
	return fmt.Sprint(findMarker(input, length))
}
