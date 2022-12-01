package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day01 struct{}

func (Day01) Name() string {
	return "day01"
}

func blockSum(block string) (int, error) {
	numbers := strings.Split(block, "\n")
	total := 0
	for _, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			return 0, err
		}
		total += n
	}
	return total, nil
}

func (Day01) Solution(input string, part2 bool) string {
	blocks := strings.Split(input, "\n\n")
	sums := make([]int, 0)
	for _, block := range blocks {
		sum, err := blockSum(block)
		if err != nil {
			fmt.Printf("invalid block: %s", block)
			return ""
		}
		sums = append(sums, sum)
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	if part2 {
		return fmt.Sprint(sums[0] + sums[1] + sums[2])
	}
	return fmt.Sprint(sums[0])
}
