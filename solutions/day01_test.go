package solutions

import "testing"

func TestDay01(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		part2  bool
		output string
	}{
		{"example1", "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", false, "24000"},
		{"example2", "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", true, "45000"},
	}

	day := Day01{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if distance := day.Solution(tc.input, tc.part2); distance != tc.output {
				t.Errorf("expected %s, got %s", tc.output, distance)
			}
		})
	}
}
