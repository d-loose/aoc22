package solutions

import (
	"fmt"
	"strings"
)

type Day02 struct{}

func (Day02) Name() string {
	return "day02"
}

type shape int

const (
	rock shape = iota
	paper
	scissors
)

type outcome int

const (
	win outcome = iota
	draw
	loss
)

var shapes = map[string]shape{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var outcomes = map[string]outcome{
	"X": loss,
	"Y": draw,
	"Z": win,
}

var shapePoints = map[shape]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}
var outcomePoints = map[outcome]int{
	win:  6,
	draw: 3,
	loss: 0,
}

func match(me, opponent shape) outcome {
	if me == opponent {
		return draw
	}
	if (me+1)%3 == opponent {
		return loss
	}
	return win
}

func chooseShape(opponent shape, outcome outcome) shape {
	switch outcome {
	case win:
		return (opponent + 1) % 3
	case loss:
		return (opponent + 2) % 3
	case draw:
		fallthrough
	default:
		return opponent
	}
}

func roundScore(me, opponent shape) int {
	result := match(me, opponent)
	return shapePoints[me] + outcomePoints[result]
}

func (Day02) Solution(input string, part2 bool) string {
	total := 0
	for _, round := range strings.Split(input, "\n") {
		codes := strings.Split(round, " ")
		opponent := shapes[codes[0]]
		var me shape
		if part2 {
			me = chooseShape(opponent, outcomes[codes[1]])
		} else {
			me = shapes[codes[1]]
		}
		score := roundScore(me, opponent)
		total += score
	}
	return fmt.Sprint(total)
}
