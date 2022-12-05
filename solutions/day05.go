package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day05 struct{}

func (Day05) Name() string {
	return "day05"
}

type stack []rune

func (s *stack) pop() rune {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}
func (s *stack) push(value rune) {
	*s = append(*s, value)
}

type crane struct {
	length int
	stacks []stack
}

func (c crane) move(n, from, to int, crateMover9001 bool) {
	buffer := make(stack, 0)
	for i := 0; i < n; i++ {
		crate := c.stacks[from-1].pop()
		if crateMover9001 {
			buffer.push(crate)
		} else {
			c.stacks[to-1].push(crate)
		}
	}
	if crateMover9001 {
		for i := 0; i < n; i++ {
			c.stacks[to-1].push(buffer.pop())
		}
	}
}

func (c crane) topRow() string {
	row := make(stack, 0)
	for pos := 0; pos < c.length; pos++ {
		row.push(c.stacks[pos][len(c.stacks[pos])-1])
	}
	return string(row)
}

func (c crane) String() string {
	var sb strings.Builder

	for pos := 0; pos < c.length; pos++ {
		sb.WriteString(fmt.Sprintf("%d: %s\n", pos, string(c.stacks[pos])))
	}
	return sb.String()
}

func parseHeader(header string) crane {
	lines := strings.Split(header, "\n")
	length := (len(lines[0]) + 1) / 4
	stacks := make([]stack, length)
	for pos := 0; pos < length; pos++ {
		stacks[pos] = make(stack, 0)
	}
	for n := len(lines) - 2; n >= 0; n-- {
		for pos := 0; pos < length; pos++ {
			if crate := []rune(lines[n])[4*pos+1]; crate != ' ' {
				stacks[pos].push(crate)
			}
		}
	}
	return crane{length: length, stacks: stacks}
}

func (Day05) Solution(input string, part2 bool) string {
	sections := strings.Split(input, "\n\n")
	crane := parseHeader(sections[0])
	for _, line := range strings.Split(sections[1], "\n") {
		tokens := strings.Split(line, " ")
		n, _ := strconv.Atoi(tokens[1])
		from, _ := strconv.Atoi(tokens[3])
		to, _ := strconv.Atoi(tokens[5])
		crane.move(n, from, to, part2)
	}
	return crane.topRow()
}
