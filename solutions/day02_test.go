package solutions

import "testing"

func TestDay02(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		part2  bool
		output string
	}{
		{"example1", "A Y\nB X\nC Z", false, "15"},
		{"example2", "A Y\nB X\nC Z", true, "12"},
	}

	day := Day02{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if distance := day.Solution(tc.input, tc.part2); distance != tc.output {
				t.Errorf("expected %s, got %s", tc.output, distance)
			}
		})
	}
}
