package utils

import (
	"os"
	"path"
	"strings"
	"testing"
)

type TestCase struct {
	Name   string
	Input  string
	Part2  bool
	Output string
}

type Day interface {
	Name() string
	Solution(string, bool) string
}

func GetInput(name string) (string, error) {
	input, err := os.ReadFile(path.Join("inputs", name))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(input)), nil
}

func RunTests(t *testing.T, day Day, testCases []TestCase) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if result := day.Solution(tc.Input, tc.Part2); result != tc.Output {
				t.Errorf("expected %s, got %s", tc.Output, result)
			}
		})
	}
}
